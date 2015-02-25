package main

import (
	"github.com/nlopes/slack"
	"log"
)

type Slack struct {
	Name   string
	Client *slack.Slack
}

func NewSlack(name, token string) *Slack {
	return &Slack{
		Name:   name,
		Client: slack.New(token),
	}
}

func (s *Slack) onMessage(message *Message) {
	_, _, err := s.Client.PostMessage(
		message.Channel,
		message.Body,
		slack.PostMessageParameters{
			Username: s.Name,
		},
	)

	if err != nil {
		log.Printf("[error] Failed to send message to %s: %s\n", message.Channel, err)
	}
}
