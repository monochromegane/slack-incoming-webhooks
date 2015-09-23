package slack_incoming_webhooks

type Field struct {
	Title string `json:"title,omitempty"`
	Value string `json:"value,omitempty"`
	Short bool   `json:"short,omitempty"`
}

func (f Field) IsEmpty() bool {
	return f.Title == "" &&
		f.Value == "" &&
		!f.Short
}
