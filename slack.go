package main

import (
	"github.com/slack-go/slack"
	"regexp"
)

type Slack struct {
	Name   string
	Icon   string
	Client *slack.Client
}

func NewSlack(name, icon, token string) *Slack {
	return &Slack{
		Name:   name,
		Icon:   icon,
		Client: slack.New(token),
	}
}

func (s *Slack) onMessage(message *Message) error {
	postMessage := slack.PostMessageParameters{
		Username:  message.Name,
		LinkNames: 1,
	}

	re := regexp.MustCompile("^:.*:$")
	if re.MatchString(message.Icon) {
		postMessage.IconEmoji = message.Icon
	} else {
		postMessage.IconURL = message.Icon
	}

	var attachment slack.Attachment
	if message.Attachment != nil {
		attachment = slack.Attachment{
			Fallback:   message.Attachment.Fallback,
			Color:      message.Attachment.Color,
			Pretext:    message.Attachment.Pretext,
			AuthorName: message.Attachment.AuthorName,
			AuthorLink: message.Attachment.AuthorLink,
			AuthorIcon: message.Attachment.AuthorIcon,
			Title:      message.Attachment.Title,
			TitleLink:  message.Attachment.TitleLink,
			Text:       message.Attachment.Text,
			ImageURL:   message.Attachment.ImageURL,
			MarkdownIn: []string{"text", "pretext", "fields"},
		}
		if len(message.Attachment.Fields) > 0 {
			fields := make([]slack.AttachmentField, len(message.Attachment.Fields))
			for i := range fields {
				fields[i].Title = message.Attachment.Fields[i].Title
				fields[i].Value = message.Attachment.Fields[i].Value
				fields[i].Short = message.Attachment.Fields[i].Short
			}
			attachment.Fields = fields
		}
	}

	_, _, err := s.Client.PostMessage(
		message.Channel,
		slack.MsgOptionText(message.Message, false),
		slack.MsgOptionAttachments(attachment),
		slack.MsgOptionPostMessageParameters(postMessage),
	)

	return err
}
