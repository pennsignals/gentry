package main

import (
	"bytes"
	"testing"
)

func TestPostMessageEncoder(t *testing.T) {
	buffer := bytes.NewBuffer([]byte(`
		[
			{
				"Node": "c2409d674e5b",
				"CheckID": "service:web",
				"Name": "Service 'web' check",
				"Status": "critical",
				"Notes": "",
				"Output": "dial tcp 127.0.0.1:8080: getsockopt: connection refused",
				"ServiceID": "web",
				"ServiceName": "web",
				"ServiceTags": ["primary"],
				"CreateIndex": 6,
				"ModifyIndex": 11
			},
			{
				"Node": "c2409d674e5b",
				"CheckID": "serfHealth",
				"Name": "Serf Health Status",
				"Status": "passing",
				"Notes": "",
				"Output": "Agent alive and reachable",
				"ServiceID": "",
				"ServiceName": "",
				"ServiceTags": [],
				"CreateIndex": 5,
				"ModifyIndex": 5
			}
		]
	`))
	encoder := NewPostMessageEncoder()
	reader, err := NewBufferedReader(buffer, encoder)
	if err != nil {
		t.Error(err)
	}
	if err := reader.Parse(); err != nil {
		t.Error(err)
	}
	// Type assertion: https://tour.golang.org/methods/15
	message := encoder.Product().(*PostMessage)
	var actual, expected interface{}
	actual, expected = len(message.Attachments), 2
	if expected != actual {
		t.Errorf("main: expected %d, got %d instead", expected, actual)
	}
	actual, expected = len(message.Attachments[0].Fields), 1
	if expected != actual {
		t.Errorf("main: expected %d, got %d instead", expected, actual)
	}
	actual, expected = message.Attachments[0].Color, ColorDanger
	if expected != actual {
		t.Errorf("main: expected %s, got %s instead", expected, actual)
	}
	actual, expected = message.Attachments[1].Color, ColorGood
	if expected != actual {
		t.Errorf("main: expected %s, got %s instead", expected, actual)
	}
}
