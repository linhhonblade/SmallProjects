package pubsub

import (
	"context"
	sctx "github.com/linhhonblade/service-context"
	log "github.com/sirupsen/logrus"
	"social_todo_app_go/common"
	"sync"
)

type localPubSub struct {
	name         string
	messageQueue chan *Message
	mapChannel   map[string][]chan *Message
	locker       *sync.RWMutex
}

func (ps *localPubSub) ID() string {
	return ps.name
}

func (ps *localPubSub) InitFlags() {
}

func (ps *localPubSub) Activate(serviceContext sctx.ServiceContext) error {
	return nil
}

func (ps *localPubSub) Stop() error {
	return nil
}

func NewLocalPubSub(name string) *localPubSub {
	pb := &localPubSub{
		name:         name,
		messageQueue: make(chan *Message, 10000),
		mapChannel:   make(map[string][]chan *Message),
		locker:       new(sync.RWMutex),
	}
	pb.run()
	return pb
}

func (ps *localPubSub) Publish(ctx context.Context, topic string, data *Message) error {
	data.SetChannel(topic)
	go func() {
		defer common.Recover()
		ps.messageQueue <- data
		log.Println("New message published:", data.String())
	}()
	return nil
}

func (ps *localPubSub) Subscribe(ctx context.Context, topic string) (ch <-chan *Message, unsubscribe func()) {
	c := make(chan *Message)

	// can thiệp vào ps.mapChannel
	ps.locker.Lock()
	if val, ok := ps.mapChannel[topic]; ok {
		val = append(ps.mapChannel[topic], c)
		ps.mapChannel[topic] = val
	} else {
		ps.mapChannel[topic] = []chan *Message{c}
	}
	ps.locker.Unlock()

	return c, func() {
		log.Println("Unsubscribe")
		if chans, ok := ps.mapChannel[topic]; ok {
			for i := range chans {
				if chans[i] == c {
					chans = append(chans[:i], chans[i+1:]...)
					ps.locker.Lock()
					ps.mapChannel[topic] = chans
					ps.locker.Unlock()

					close(c)
					break
				}
			}
		}
	}
}

func (ps *localPubSub) run() error {
	go func() {
		defer common.Recover()
		for {
			mess := <-ps.messageQueue
			log.Println("Message queue:", mess.String())
			if subs, ok := ps.mapChannel[mess.channel]; ok {
				for i := range subs {
					go func(c chan *Message) {
						defer common.Recover()
						c <- mess
					}(subs[i])
				}
			}
		}
	}()
	return nil
}
