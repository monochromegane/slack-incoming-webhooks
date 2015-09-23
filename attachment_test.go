package slack_incoming_webhooks

import "testing"

func TestAddField(t *testing.T) {
	attachment := Attachment{}
	field := &Field{}
	attachment.AddField(field)

	if len(attachment.Fields) > 0 {
		t.Errorf("Attachment should not add a empty field, but field size is %b.", len(attachment.Fields))
	}

	field.Title = "title"
	attachment.AddField(field)

	if len(attachment.Fields) != 1 {
		t.Errorf("Attachment should add a field, but field size is %b.", len(attachment.Fields))
	}
}
