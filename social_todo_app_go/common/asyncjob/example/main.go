package main

import (
	"context"
	"log"
	"social_todo_app_go/common/asyncjob"
	"time"
)

func main() {
	j1 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("init job 1")
		return nil
	},
		asyncjob.WithName("Job 1"),
		asyncjob.WithRetryDuration([]time.Duration{time.Second * 5}),
	)
	j2 := asyncjob.NewJob(func(ctx context.Context) error {
		log.Println("init job 2")
		return nil
	},
		asyncjob.WithName("Job 2"),
		asyncjob.WithRetryDuration([]time.Duration{time.Second * 5}),
	)
	jm := asyncjob.NewGroup(true, j1, j2)
	if err := jm.Run(context.Background()); err != nil {
		log.Println(err)
	}
}
