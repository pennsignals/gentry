package main

import (
	"encoding/json"
	"io"
)

const DefaultBufferSize int = 4096

type Encoder interface {
	Encode(check *Check)
	Product() Product
}

type Product interface{}

type BufferedReader struct {
	bytes   []byte
	encoder Encoder
}

func NewBufferedReader(reader io.Reader, encoder Encoder) (*BufferedReader, error) {
	bytes := make([]byte, DefaultBufferSize)
	n, err := reader.Read(bytes)
	if err != nil {
		return nil, err
	}
	return &BufferedReader{bytes[:n], encoder}, nil
}

func (r *BufferedReader) Parse() error {
	var checks Checks
	if err := json.Unmarshal(r.bytes, &checks); err != nil {
		return err
	}
	for _, check := range checks {
		r.encoder.Encode(check)
	}
	return nil
}

type PostMessageEncoder struct {
	message *PostMessage
}

func NewPostMessageEncoder() *PostMessageEncoder {
	return &PostMessageEncoder{new(PostMessage)}
}

// Encode converts a Check object into an Attachment object and
// appends it to the end of the Attachments slice.
func (c *PostMessageEncoder) Encode(check *Check) {
	colors := map[State]Color{
		StateCritical: ColorDanger,
		StatePassing:  ColorGood,
		StateWarning:  ColorWarning,
	}
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

// Product returns a PostMessage object.
func (c *PostMessageEncoder) Product() Product {
	return c.message
}
