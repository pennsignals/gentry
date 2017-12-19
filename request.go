package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type RequestCreator interface {
	Create() *http.Request
}

type PostMessageRequestCreator struct {
	address *url.URL
	message *PostMessage
}

func NewPostMessageRequestCreator(address string, message *PostMessage) (*PostMessageRequestCreator, error) {
	identifier, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	return &PostMessageRequestCreator{identifier, message}, nil
}

func (c *PostMessageRequestCreator) Create() *http.Request {
	body, _ := json.Marshal(c.message)
	request, _ := http.NewRequest(http.MethodPost, c.address.String(), bytes.NewReader(body))
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.message.Token))
	request.Header.Add("Content-Type", "application/json")
	return request
}

type Requester struct {
	client  *http.Client
	creator RequestCreator
}

func NewRequester(creator RequestCreator, options ...func(client *http.Client)) *Requester {
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
