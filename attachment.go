package slack_incoming_webhooks

type Attachment struct {
	Fallback   string   `json:"fallback,omitempty"`
	Color      string   `json:"color,omitempty"`
	Pretext    string   `json:"pretext,omitempty"`
	AuthorName string   `json:"author_name,omitempty"`
	AuthorLink string   `json:"author_link,omitempty"`
	AuthorIcon string   `json:"author_icon,omitempty"`
	Title      string   `json:"title,omitempty"`
	TitleLink  string   `json:"title_link,omitempty"`
	Text       string   `json:"text,omitempty"`
	Fields     []*Field `json:"fields,omitempty"`
	ImageURL   string   `json:"image_url,omitempty"`
	ThumbURL   string   `json:"thumb_url,omitempty"`
	Footer     string   `json:"footer,omitempty"`
	FooterIcon string   `json:"footer_icon,omitempty"`
	Timestamp  int64    `json:"ts,omitempty"`
	MarkdownIn []string `json:"mrkdwn_in,omitempty"`
}

func (a Attachment) IsEmpty() bool {
	return a.Fallback == "" &&
		a.Color == "" &&
		a.Pretext == "" &&
		a.AuthorName == "" &&
		a.AuthorLink == "" &&
		a.AuthorIcon == "" &&
		a.Title == "" &&
		a.TitleLink == "" &&
		a.Text == "" &&
		len(a.Fields) == 0 &&
		a.ImageURL == "" &&
		a.ThumbURL == "" &&
		len(a.MarkdownIn) == 0
}

func (a *Attachment) AddField(f *Field) {
	if !f.IsEmpty() {
		a.Fields = append(a.Fields, f)
	}
}
