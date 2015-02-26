package main

type Message struct {
	Group  string
	Body   string
	Result chan error
}

func NewMessage(group, body string, ch chan error) *Message {
	return &Message{
		Group:  group,
		Body:   body,
		Result: ch,
	}
}
