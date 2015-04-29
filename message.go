package main

type Message struct {
	Channel    string
	Message    string
	Name       string
	Icon       string
	Attachment *Attachment
	Manual     bool
	Result     chan error
}

type Attachment struct {
	Color      string
	Pretext    string
	AuthorName string
	AuthorLink string
	AuthorIcon string
	Title      string
	TitleLink  string
	Fallback   string
	Text       string
	Fields     []Field
	ImageURL   string
}

type Field struct {
	Title string
	Value string
	Short bool
}

func (p *Param) HasAttachment() bool {
	if (p.Text != "") ||
		(p.Color != "") ||
		(p.Fallback != "") ||
		(p.Pretext != "") ||
		(p.AuthorName != "") ||
		(p.AuthorLink != "") ||
		(p.AuthorIcon != "") ||
		(p.Title != "") ||
		(p.TitleLink != "") ||
		(p.ImageURL != "") ||
		(len(p.FieldTitle) > 0) ||
		(len(p.FieldValue) > 0) {
		return true
	}

	return false
}

func (p *Param) HasField() bool {
	if (len(p.FieldTitle) > 0) || (len(p.FieldValue) > 0) {
		return true
	}

	return false
}

func (p *Param) Adjust() {
	if p.Name == "" {
		p.Name = name
	}
	if p.Icon == "" {
		p.Icon = icon
	}

	if p.Manual == false {
		if p.Message != "" && p.Text == "" && p.Color != "" {
			p.Text = p.Message
			p.Message = ""
		}
		if p.Text != "" && p.Fallback == "" {
			p.Fallback = p.Text
		}
	}
}

func NewMessage(p Param, ch chan error) *Message {
	p.Adjust()

	message := Message{
		Channel: p.Channel,
		Message: p.Message,
		Name:    p.Name,
		Icon:    p.Icon,
		Manual:  p.Manual,
		Result:  ch,
	}

	if p.HasAttachment() {
		message.Attachment = NewAttachment(p)
	}

	return &message
}

func NewAttachment(p Param) *Attachment {
	attachment := Attachment{
		Fallback:   p.Fallback,
		Color:      p.Color,
		Pretext:    p.Pretext,
		AuthorName: p.AuthorName,
		AuthorLink: p.AuthorLink,
		AuthorIcon: p.AuthorIcon,
		Title:      p.Title,
		TitleLink:  p.TitleLink,
		Text:       p.Text,
		ImageURL:   p.ImageURL,
	}

	if p.HasField() {
		attachment.Fields = NewFields(p)
	}

	return &attachment
}

func NewFields(p Param) []Field {
	field_title_max := len(p.FieldTitle)
	field_value_max := len(p.FieldValue)
	field_short_max := len(p.FieldShort)

	field_max := 0
	if field_title_max >= field_value_max {
		field_max = field_title_max
	} else {
		field_max = field_value_max
	}

	fields := make([]Field, field_max)
	for i := range fields {
		if i < field_title_max {
			fields[i].Title = p.FieldTitle[i]
		}
		if i < field_value_max {
			fields[i].Value = p.FieldValue[i]
		}
		if i < field_short_max {
			fields[i].Short = p.FieldShort[i]
		}
	}

	return fields
}
