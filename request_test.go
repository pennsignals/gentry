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

func (t *MockRoundTripper) Recorder() *httptest.ResponseRecorder {
	return t.recorder
}

func (t *MockRoundTripper) RoundTrip(request *http.Request) (*http.Response, error) {
	t.recorder.HeaderMap = make(http.Header)
	t.recorder.HeaderMap.Set("Content-Type", "application/json")
	io.WriteString(t.recorder, `
		{
			"ok": false,
			"error": "not_authed"
		}
	`)
	return t.recorder.Result(), nil
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
	creator, err := NewPostMessageRequestCreator(&PostMessage{Token: "xoxb-"})
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

func TestRequesterMockRequest(t *testing.T) {
	creator, err := NewPostMessageRequestCreator(new(PostMessage))
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
