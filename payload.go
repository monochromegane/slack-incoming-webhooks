package slack_incoming_webhooks

type Payload struct {
	Channel     string        `json:"channel,omitempty"`
	Username    string        `json:"username,omitempty"`
	Text        string        `json:"text,omitempty"`
	IconEmoji   string        `json:"icon_emoji,omitempty"`
	IconURL     string        `json:"icon_url,omitempty"`
	Attachments []*Attachment `json:"attachments,omitempty"`
}

func (p *Payload) AddAttachment(a *Attachment) {
	if !a.IsEmpty() {
		p.Attachments = append(p.Attachments, a)
	}
}
