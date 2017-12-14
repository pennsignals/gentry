package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type Creator interface {
	Create() *http.Request
}

type PostMessageCreator struct {
	address *url.URL
	message *PostMessage
}

func NewPostMessageCreator(address string, message *PostMessage) (*PostMessageCreator, error) {
	identifier, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	return &PostMessageCreator{identifier, message}, nil
}

func (c *PostMessageCreator) Create() *http.Request {
	body, _ := json.Marshal(c.message)
	request, _ := http.NewRequest(http.MethodPost, c.address.String(), bytes.NewReader(body))
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.message.Token))
	request.Header.Add("Content-Type", "application/json")
	return request
}

type Requester struct {
	client  *http.Client
	creator Creator
}

func NewRequester(creator Creator, options ...func(client *http.Client)) *Requester {
	var client http.Client
	for _, option := range options {
		option(&client)
	}
	return &Requester{&client, creator}
}

func (r *Requester) Client() *http.Client {
	return r.client
}

func (r *Requester) Request() (*http.Response, error) {
	request := r.creator.Create()
	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
