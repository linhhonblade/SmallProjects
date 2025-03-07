package asyncjob

import (
	"context"
	"log"
	"social_todo_app_go/common"
	"sync"
)

type group struct {
	jobs         []Job
	isConcurrent bool
	wg           *sync.WaitGroup
}

func NewGroup(isConcurrent bool, jobs ...Job) *group {
	return &group{
		jobs:         jobs,
		isConcurrent: isConcurrent,
		wg:           &sync.WaitGroup{},
	}
}
func (g *group) Run(ctx context.Context) error {
	g.wg.Add(len(g.jobs))
	errChan := make(chan error, len(g.jobs))
	for i := range g.jobs {
		if g.isConcurrent {
			go func(aj Job) {
				defer common.Recover()
				errChan <- g.runJob(ctx, aj)
				g.wg.Done()
			}(g.jobs[i])
			continue
		}
		job := g.jobs[i]

		// COR design pattern
		err := g.runJob(ctx, job)
		if err != nil {
			return err
		}
		errChan <- err
		g.wg.Done()
	}
	g.wg.Wait()
	var err error
	for i := 1; i <= len(g.jobs); i++ {
		if v := <-errChan; v != nil {
			return v
		}
	}
	return err
}

func (g *group) runJob(ctx context.Context, j Job) error {
	err := j.Execute(ctx)
	if err != nil {
		for {
			log.Println(err)
			if j.State() == StateRetryFailed {
				return err
			}
			if j.Retry(ctx) == nil {
				return nil
			}
		}
	}
	return nil
}
