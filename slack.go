package main

import (
	"github.com/nlopes/slack"
)

type Slack struct {
	Name   string
	Icon   string
	Client *slack.Slack
}

func NewSlack(name, icon, token string) *Slack {
	return &Slack{
		Name:   name,
		Icon:   icon,
		Client: slack.New(token),
	}
}

func (s *Slack) onMessage(message *Message) error {
	_, _, err := s.Client.PostMessage(
		message.Group,
		message.Body,
		slack.PostMessageParameters{
			Username:  s.Name,
			IconURL:   s.Icon,
			LinkNames: 1,
		},
	)

	return err
}
