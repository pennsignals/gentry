package main

import (
	"encoding/json"
	"fmt"
	"io"
)

const DefaultBufferSize int = 4096

type Encoder interface {
	Encode(checks Checks)
}

type bufferedReader struct {
	buffer  []byte
	encoder Encoder
}

func NewBufferedReader(reader io.Reader, encoder Encoder) (*bufferedReader, error) {
	buffer := make([]byte, DefaultBufferSize)
	n, err := reader.Read(buffer)
	if err != nil {
		return nil, err
	}
	return &bufferedReader{buffer[:n], encoder}, nil
}

func (r *bufferedReader) Parse() error {
	var checks Checks
	if err := json.Unmarshal(r.buffer, &checks); err != nil {
		return err
	}
	r.encoder.Encode(checks)
	return nil
}

type postMessageEncoder struct {
	message *PostMessage
}

// NewPostMessageEncoder returns a new postMessageEncoder given a
// channel and a token.
func NewPostMessageEncoder(channel, token string) *postMessageEncoder {
	return &postMessageEncoder{&PostMessage{
		Channel: channel,
		Token:   token,
	}}
}

// Encode converts a Checks object into a PostMessage object.
func (c *postMessageEncoder) Encode(checks Checks) {
	c.message.Text = fmt.Sprintf("Consul catalog contains %d registered nodes", len(checks))
	// Convert each Check object into an Attachment object.
	colors := map[State]Color{
		StateCritical: ColorDanger,
		StatePassing:  ColorGood,
		StateWarning:  ColorWarning,
	}
	for _, check := range checks {
		fields := make(Fields, len(check.ServiceTags))
		for i, tag := range check.ServiceTags {
			fields[i] = &Field{Short: true, Title: "Service Tag", Value: tag}
		}
		c.message.Attachments.Add(&Attachment{
			Color:  colors[check.Status],
			Fields: fields,
			Text:   check.Output,
			Title:  check.Name,
		})
	}
}

// Product returns a PostMessage object.
func (c *postMessageEncoder) Product() *PostMessage {
	return c.message
}
