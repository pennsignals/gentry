package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type RequestFactory interface {
	Create() *http.Request
}

type PostMessageRequestFactory struct {
	address *url.URL
	message *PostMessage
}

func NewPostMessageRequestFactory(address string, message *PostMessage) (*PostMessageRequestFactory, error) {
	identifier, err := url.Parse(address)
	if err != nil {
		return nil, err
	}
	return &PostMessageRequestFactory{identifier, message}, nil
}

func (c *PostMessageRequestFactory) Create() *http.Request {
	body, _ := json.Marshal(c.message)
	request, _ := http.NewRequest(http.MethodPost, c.address.String(), bytes.NewReader(body))
	request.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.message.Token))
	request.Header.Add("Content-Type", "application/json")
	return request
}

type Requester struct {
	client  *http.Client
	factory RequestFactory
}

func NewRequester(factory RequestFactory, options ...func(client *http.Client)) *Requester {
	var client http.Client
	for _, option := range options {
		option(&client)
	}
	return &Requester{&client, factory}
}

func (r *Requester) Client() *http.Client {
	return r.client
}

func (r *Requester) Request() (*http.Response, error) {
	request := r.factory.Create()
	response, err := r.client.Do(request)
	if err != nil {
		return nil, err
	}
	return response, nil
}
