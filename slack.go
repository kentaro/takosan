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
	postMessage := slack.PostMessageParameters{
		LinkNames: 1,
	}

	// Set bot name
	if message.Name != "" {
		postMessage.Username = message.Name
	} else {
		postMessage.Username = s.Name
	}

	// Set bot Icon
	iconString := ""
	if message.Icon != "" {
		iconString = message.Icon
	} else {
		iconString = s.Icon
	}

	// Switch IconURL or IconEmoji
	re := regexp.MustCompile("^:.*:$")
	if re.MatchString(iconString) {
		postMessage.IconEmoji = iconString
	} else {
		postMessage.IconURL = iconString
	}

	messageText := message.Body
	attachment := slack.Attachment{}

	attachment.Color = message.Color
	attachment.Pretext = message.Pretext
	attachment.AuthorName = message.AuthorName
	attachment.AuthorLink = message.AuthorLink
	attachment.AuthorIcon = message.AuthorIcon
	attachment.Text = message.Text
	attachment.Fallback = message.Text
	attachment.Title = message.Title
	attachment.TitleLink = message.TitleLink
	attachment.ImageURL = message.ImageURL

	// Automatic actions
	if message.Manual == false {
		if attachment.Color != "" && message.Body != "" && message.Text == "" {
			messageText = ""
			attachment.Text = message.Body
			attachment.Fallback = message.Body
		}
	}

	postMessage.Attachments = []slack.Attachment{attachment}
	postMessage.Parse = message.Parse

	_, _, err := s.Client.PostMessage(message.Group, messageText, postMessage)
	return err
}
