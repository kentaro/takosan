package main

import (
	"github.com/nlopes/slack"
	"log"
)

type Slack struct {
	Client *slack.Slack
}

func NewSlack(token string) *Slack {
	return &Slack{
		Client: slack.New(token),
	}
}

func (s *Slack) onMessage(message *Message) {
	_, _, err := s.Client.PostMessage(
		message.Channel,
		message.Body,
		slack.PostMessageParameters{},
	)

	if err != nil {
		log.Printf("[error] Failed to send message to %s: %s\n", message.Channel, err)
	}
}
