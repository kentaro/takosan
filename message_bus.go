package main

type Bus struct {
	queue chan *Message
}

type Subscriber interface {
	onMessage(message *Message)
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
			subscriber.onMessage(message)
		}
	}()
}
