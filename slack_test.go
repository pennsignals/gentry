package main

import (
	"encoding/json"
	"testing"
)

func TestTypeAttachmentEncode(t *testing.T) {
	var attachments Attachments
	attachments.Add(&Attachment{Color: ColorDanger})
	message := &PostMessage{Attachments: attachments}
	b, err := json.Marshal(message)
	if err != nil {
		t.Error(err)
	}
	if err := json.Unmarshal(b, &message); err != nil {
		t.Error(err)
	}
	actual, expected := message.Attachments[0].Color.String(), "danger"
	if expected != actual {
		t.Errorf("main: expected %s, got %s", expected, actual)
	}
}
