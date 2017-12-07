package main

import (
	"encoding/json"
	"io"
)

const DefaultBufferSize int = 4096

type Converter interface {
	Convert(check *Check)
	Product() Product
}

type Product interface{}

type BufferedReader struct {
	bytes     []byte
	converter Converter
}

func NewBufferedReader(reader io.Reader, converter Converter) (*BufferedReader, error) {
	bytes := make([]byte, DefaultBufferSize)
	n, err := reader.Read(bytes)
	if err != nil {
		return nil, err
	}
	return &BufferedReader{bytes[:n], converter}, nil
}

func (r *BufferedReader) Parse() error {
	var checks Checks
	if err := json.Unmarshal(r.bytes, &checks); err != nil {
		return err
	}
	for _, check := range checks {
		r.converter.Convert(check)
	}
	return nil
}

type PostMessageConverter struct {
	message *PostMessage
}

func NewPostMessageConverter() *PostMessageConverter {
	return &PostMessageConverter{new(PostMessage)}
}

// Convert converts a Check object into an Attachment object and
// appends it to the end of the Attachments slice.
func (c *PostMessageConverter) Convert(check *Check) {
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
func (c *PostMessageConverter) Product() Product {
	return c.message
}
