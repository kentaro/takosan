package main

type Message struct {
	Group      string
	Body       string
	Name       string
	Icon       string
	Color      string
	Pretext    string
	AuthorName string
	AuthorLink string
	AuthorIcon string
	Title      string
	TitleLink  string
	Text       string
	ImageURL   string
	Parse      string
	Manual     bool
	Result     chan error
}

func NewMessage(group, body string, ch chan error) *Message {
	return &Message{
		Group:  group,
		Body:   body,
		Result: ch,
	}
}
