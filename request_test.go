package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type MockRoundTripper struct {
	recorder *httptest.ResponseRecorder
}

func NewMockRoundTripper() *MockRoundTripper {
	return &MockRoundTripper{new(httptest.ResponseRecorder)}
}

func (rt *MockRoundTripper) Recorder() *httptest.ResponseRecorder {
	return rt.recorder
}

func (rt *MockRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	rt.recorder.HeaderMap = make(http.Header)
	rt.recorder.HeaderMap.Set("Content-Type", "application/json")
	io.WriteString(rt.recorder, `
		{
			"ok": false,
			"error": "not_authed"
		}
	`)
	return rt.recorder.Result(), nil
}

func TestNewRequesterClientConfiguration(t *testing.T) {
	requester := NewRequester(nil, func(client *http.Client) {
		client.Timeout = time.Minute
	})
	client := requester.Client()
	actual, expected := client.Timeout.String(), "1m0s"
	if expected != actual {
		t.Errorf("main: expected %q, got %q instead", expected, actual)
	}
}

func TestPostMessageCreator(t *testing.T) {
	creator, err := NewPostMessageCreator("", &PostMessage{Token: "xoxb-"})
	if err != nil {
		t.Error(err)
	}
	request := creator.Create()
	var actual, expected interface{}
	actual, expected = request.Header.Get("Content-Type"), "application/json"
	if expected != actual {
		t.Errorf("main: expected %q, got %q instead", expected, actual)
	}
	// PostMessageCreator uses the PostMessage Token for the
	// Authorization header.
	actual, expected = request.Header.Get("Authorization"), "Bearer xoxb-"
	if expected != actual {
		t.Errorf("main: expected %q, got %q instead", expected, actual)
	}
}

func TestPostMessageCreatorMalformedAddress(t *testing.T) {
	if _, err := NewPostMessageCreator(":", nil); err == nil {
		t.Errorf("main: expected NewPostMessageCreator to return an error value")
	}
}

func TestRequesterMockRequest(t *testing.T) {
	address := "https://slack.com/api/api.test"
	creator, err := NewPostMessageCreator(address, new(PostMessage))
	if err != nil {
		t.Error(err)
	}
	transport := NewMockRoundTripper()
	requester := NewRequester(creator, func(client *http.Client) {
		client.Transport = transport
	})
	response, err := requester.Request()
	if err != nil {
		t.Error(err)
	}
	var actual, expected interface{}
	actual, expected = response.StatusCode, 200
	if expected != actual {
		t.Errorf("main: expected %d, got %d instead", expected, actual)
	}
	actual, expected = response.Header.Get("Content-Type"), "application/json"
	if expected != actual {
		t.Errorf("main: expected %q, got %q instead", expected, actual)
	}
}
