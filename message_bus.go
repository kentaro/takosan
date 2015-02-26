package main

import (
	"time"
)

type Bus struct {
	queue chan *Message
}

type Subscriber interface {
	onMessage(message *Message) error
}

var MessageBus = &Bus{
	queue: make(chan *Message),
}

func (b Bus) Publish(message *Message) {
	b.queue <- message
}

func (b Bus) Subscribe(subscriber Subscriber) {
	go func() {
		for {
			message := <-b.queue

			// To comply with API rate limit requirement
			done := make(chan interface{}, 1)
			go func() {
				done <- time.After(1 * time.Second)
			}()

			err := subscriber.onMessage(message)
			message.Result <- err

			<-done
		}
	}()
}
