package main

type Message struct {
	Channel string
	Body    string
}

func NewMessage(channel, body string) *Message {
	return &Message{
		Channel: channel,
		Body:    body,
	}
}
