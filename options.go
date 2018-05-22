package ethrpc

import (
	"io"
	"net/http"
)

type httpClient interface {
	Post(url string, contentType string, body io.Reader) (*http.Response, error)
}

type logger interface {
	Println(v ...interface{})
}

// WithHTTPClient set custom http client
func WithHTTPClient(clt httpClient) func(client *Client) {
	return func(client *Client) {
		client.httpClient = clt
	}
}

// WithLogger set custom logger
func WithLogger(l logger) func(client *Client) {
	return func(client *Client) {
		client.log = l
	}
}

// WithDebug set debug flag
func WithDebug(enabled bool) func(client *Client) {
	return func(client *Client) {
		client.Debug = enabled
	}
}
