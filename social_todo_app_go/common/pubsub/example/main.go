package main

import (
	"context"
	"fmt"
	"log"
	"social_todo_app_go/common/pubsub"
	"time"
)

func main() {
	var broker pubsub.PubSub = pubsub.NewLocalPubSub("local-pubsub")
	const topic = "test-topic"
	ch1, close1 := broker.Subscribe(context.Background(), topic)
	ch2, _ := broker.Subscribe(context.Background(), topic)
	go func() {
		for v := range ch1 {
			fmt.Println("ch1", v)
		}
		log.Println("ch1 closed")
	}()
	go func() {
		for v := range ch2 {
			fmt.Println("ch2", v)
		}
		log.Println("ch2 closed")
	}()
	broker.Publish(context.Background(), topic, pubsub.NewMessage(map[string]interface{}{"key1": "value1"}))
	time.Sleep(time.Second * 3)
	close1()
	broker.Publish(context.Background(), topic, pubsub.NewMessage(map[string]interface{}{"key2": "value2"}))
	time.Sleep(time.Second)
}
