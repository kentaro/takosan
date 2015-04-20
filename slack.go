package main

import (
	"github.com/nlopes/slack"
	"regexp"
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
	fields := make([]slack.AttachmentField, len(message.Attachment.Fields))
	for i := range fields {
		fields[i].Title = message.Attachment.Fields[i].Title
		fields[i].Value = message.Attachment.Fields[i].Value
		fields[i].Short = message.Attachment.Fields[i].Short
	}

	attachment := slack.Attachment{
		Fallback:   message.Attachment.Fallback,
		Color:      message.Attachment.Color,
		Pretext:    message.Attachment.Pretext,
		AuthorName: message.Attachment.AuthorName,
		AuthorLink: message.Attachment.AuthorLink,
		AuthorIcon: message.Attachment.AuthorIcon,
		Title:      message.Attachment.Title,
		TitleLink:  message.Attachment.TitleLink,
		Text:       message.Attachment.Text,
		Fields:     fields,
		ImageURL:   message.Attachment.ImageURL,
		// MarkdownIn: []string{"text", "pretext", "fields"},
	}

	postMessage := slack.PostMessageParameters{
		Username:    message.Name,
		Attachments: []slack.Attachment{attachment},
		LinkNames:   1,
	}

	re := regexp.MustCompile("^:.*:$")
	if re.MatchString(message.Icon) {
		postMessage.IconEmoji = message.Icon
	} else {
		postMessage.IconURL = message.Icon
	}

	_, _, err := s.Client.PostMessage(message.Channel, message.Message, postMessage)
	return err
}
