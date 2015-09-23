package slack_incoming_webhooks

import "testing"

func TestAddAttachment(t *testing.T) {
	payload := Payload{}
	attachment := &Attachment{}
	payload.AddAttachment(attachment)

	if len(payload.Attachments) > 0 {
		t.Errorf("Payload should not add a empty attachment, but attachment size is %b.", len(payload.Attachments))
	}

	attachment.Text = "text"
	payload.AddAttachment(attachment)

	if len(payload.Attachments) != 1 {
		t.Errorf("Payload should add a attachment, but attachment size is %b.", len(payload.Attachments))
	}
}
