package worker_pool_example_go

import (
	"fmt"
	"log"
	"sync"
	"testing"
	"time"
)

type Pool interface {
	Start()
	Stop()
	AddWork(Task)
}

type Task interface {
	Execute() error
	OnFailure(error)
}

type SimplePool struct {
	numsWorker int
	tasks      chan Task
	start      sync.Once
	end        sync.Once
	quit       chan struct{}
}

func (p *SimplePool) Start() {
	p.start.Do(func() {
		log.Println("starting simple worker pool")
		p.startWorkers()
	})
}

func (p *SimplePool) Stop() {
	p.end.Do(func() {
		log.Println("stopping simple worker pool")
		close(p.quit)
	})
}

func (p *SimplePool) AddWork(t Task) {
	select {
	case p.tasks <- t:
	case <-p.quit:
	}
}

func (p *SimplePool) AddWorkNonBlocking(t Task) {
	go p.AddWork(t)
}

func (p *SimplePool) startWorkers() {
	for i := 0; i < p.numsWorker; i++ {
		go func(workerNum int) {
			log.Printf("starting worker %d", workerNum)

			for {
				select {
				case <-p.quit:
					log.Printf("stopping worker %d with quit channel\n", workerNum)
					return
				case task, ok := <-p.tasks:
					if !ok {
						log.Printf("stopping worker %d with closed task channel\n", workerNum)
						return
					}
					if err := task.Execute(); err != nil {
						task.OnFailure(err)
					}
				}
			}
		}(i)
	}
}

var ErrNoWorkers = fmt.Errorf("attempting to create worker pool with less than one worker")
var ErrNegativeChannelSize = fmt.Errorf("attempting to create worker pool with a negative channel size.")

func NewSimplePool(numsWorker int, channelSize int) (Pool, error) {
	if numsWorker <= 0 {
		return nil, ErrNoWorkers
	}
	if channelSize < 0 {
		return nil, ErrNegativeChannelSize
	}
	tasks := make(chan Task, channelSize)
	return &SimplePool{
		numsWorker: numsWorker,
		tasks:      tasks,
		start:      sync.Once{},
		end:        sync.Once{},
		quit:       make(chan struct{}),
	}, nil
}

func (p *SimplePool) startWorker() {
	for i := 0; i < p.numsWorker; i++ {
		go func(workerNum int) {
			log.Println("starting worker ", workerNum)
			for {
				select {
				case <-p.quit:
					log.Println("stopping worker ", workerNum, " with quit channel")
					return
				case task, ok := <-p.tasks:
					if !ok {
						log.Printf("stopping worker %d with closed tasks channel\n", workerNum)
						return
					}
					if err := task.Execute(); err != nil {
						task.OnFailure(err)
					}
				}
			}
		}(i)
	}
}

func TestWorkerPool_NewPool(t *testing.T) {
	if _, err := NewSimplePool(0, 0); err != ErrNoWorkers {
		t.Fatalf("expected error when creating pool with 0 worker, got: %v", err)
	}
	if _, err := NewSimplePool(-1, 0); err != ErrNoWorkers {
		t.Fatalf("expected error when creating pool with negative worker, got: %v", err)
	}
	if _, err := NewSimplePool(1, -1); err != ErrNegativeChannelSize {
		t.Fatalf("expected error when creating pool with negative channel size, got: %v", err)
	}
	p, err := NewSimplePool(5, 0)
	if err != nil {
		t.Fatalf("expected no error when creating pool, got: %v", err)
	}
	if p == nil {
		t.Fatal("NewSimplePool returned nil Pool for valid input")
	}
}

func TestWorkerPool_MultipleStartStopNoPanic(t *testing.T) {
	p, err := NewSimplePool(5, 0)
	if err != nil {
		t.Fatal("error creating pool: ", err)
	}
	p.Start()
	p.Stop()
	p.Start()
	p.Stop()
}

type testTask struct {
	executeFunc    func() error
	shouldErr      bool
	wg             *sync.WaitGroup
	mFailure       *sync.Mutex
	failureHandled bool
}

func newTestTask(executeFunc func() error, shouldErr bool, wg *sync.WaitGroup) *testTask {
	return &testTask{
		executeFunc: executeFunc,
		shouldErr:   shouldErr,
		wg:          wg,
		mFailure:    &sync.Mutex{},
	}
}

func (t *testTask) Execute() error {
	if t.wg != nil {
		defer t.wg.Done()
	}
	if t.executeFunc != nil {
		return t.executeFunc()
	}
	// If no executeFunc provided,just wait and error if told to do so
	time.Sleep(50 * time.Millisecond)
	if t.shouldErr {
		return fmt.Errorf("planned Execute() errpr")
	}
	return nil
}

func (t *testTask) OnFailure(e error) {
	t.mFailure.Lock()
	defer t.mFailure.Unlock()
	t.failureHandled = true
}

func (t *testTask) hitFailureCase() bool {
	t.mFailure.Lock()
	defer t.mFailure.Unlock()
	return t.failureHandled
}

func TestWorkerPool_Work(t *testing.T) {
	var tasks []*testTask
	wg := &sync.WaitGroup{}
	for i := 0; i < 20; i++ {
		wg.Add(1)
		tasks = append(tasks, newTestTask(nil, false, wg))
	}
	p, err := NewSimplePool(5, len(tasks))
	if err != nil {
		t.Fatal("error making worker pool:", err)
	}
	p.Start()
	for _, j := range tasks {
		p.AddWork(j)
	}
	wg.Wait()
	for taskNum, task := range tasks {
		if task.hitFailureCase() {
			t.Fatalf("error function called on task %d when it shouldn't be", taskNum)
		}
	}
}

func TestWorkerPool_BlockedAddWorkReleaseAfterStop(t *testing.T) {
	p, err := NewSimplePool(1, 0)
	if err != nil {
		t.Fatal("error making worker pool:", err)
	}
	p.Start()
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		// the first should start processing right away, the second two should hang
		wg.Add(1)
		go func() {
			p.AddWork(newTestTask(func() error {
				time.Sleep(20 * time.Second)
				return nil
			}, false, nil))
			wg.Done()
		}()
	}
	done := make(chan struct{})
	p.Stop()
	go func() {
		// wait on our AddWork calls to complete, then signal on the done channel
		wg.Wait()
		done <- struct{}{}
	}()
	// wait until either we hit our timeout, or we're told the AddWork calls completed
	select {
	case <-time.After(1 * time.Second):
		t.Fatal("failed because still hanging on AddWork")
	case <-done:
		//this is the success case
	}
}
