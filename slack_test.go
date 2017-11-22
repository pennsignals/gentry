package main

import (
	"encoding/json"
	"testing"
)

func TestTypeAttachmentDecode(t *testing.T) {
	// Sample Attachment response from: https://api.slack.com/docs/message-attachments
	data := []byte(`
		[
			{
				"author_icon": "http://flickr.com/icons/bobby.jpg",
				"author_link": "http://flickr.com/bobby/",
				"author_name": "Bobby Tables",
				"color": "danger",
				"fallback": "Required plain-text summary of the attachment.",
				"fields": [
					{
						"short": false,
						"title": "Priority",
						"value": "High"
					}
				],
				"footer": "Slack API",
				"footer_icon": "https://platform.slack-edge.com/img/default_application_icon.png",
				"image_url": "http://my-website.com/path/to/image.jpg",
				"pretext": "Optional text that appears above the attachment block",
				"text": "Optional text that appears within the attachment",
				"thumb_url": "http://example.com/path/to/thumb.png",
				"title": "Slack API Documentation",
				"title_link": "https://api.slack.com/",
				"ts": 123456789
			}
		]
	`)
	var attachments Attachments
	if err := json.Unmarshal(data, &attachments); err != nil {
		t.Error(err)
	}
	var actual, expected string
	actual = attachments[0].Color.String()
	expected = ColorDanger.String()
	if expected != actual {
		t.Errorf("main: expected %s, got %s", expected, actual)
	}
}
