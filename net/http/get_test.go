package http

import (
	"testing"
)

var httpCodeOK = 200
var httpMessageCodeOK = "200 OK"

func TestHttpRequestSuccess(t *testing.T) {
	get := Get{
		Url: "http://www.gmail.com",
	}
	resp, err := get.Get()
	if err != nil {
		t.Error("Expected a success request")
	}

	statusCode := resp.StatusCode()
	if httpCodeOK != statusCode {
		t.Error("Expected HTTP status code 200 but received", statusCode)
	}

	messageCode := resp.StatusMessageCode()
	if httpMessageCodeOK != messageCode {
		t.Error("Expected HTTP message 200 OK but received", messageCode)
	}
}

func TestHttpRequestFail(t *testing.T) {
	get := Get{
		Url: "www.gmail.com",
	}
	_, err := get.Get()
	if err == nil {
		t.Error("Expected a fail request")
	}
}
