package main

import (
	"github.com/nlopes/slack"
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

func (s *Slack) onMessage(message *Message) error {
	_, _, err := s.Client.PostMessage(
		message.Group,
		message.Body,
		slack.PostMessageParameters{
			Username: s.Name,
		},
	)

	return err
}
