package http

import (
	"io"
	"net/http"
)

type (
	// Request responsible to handle HTTP incoming calls.
	Request struct {
		request *http.Request
	}
)

// NewRequest creates a new instance of Request struct.
func NewRequest(r *http.Request) *Request {
	return &Request{
		request: r,
	}
}

// Body returns the current body of this request.
func (r *Request) Body() io.ReadCloser {
	return r.request.Body
}

// HasHeader verifies if the given key exists in the headers.
func (r *Request) HasHeader(k string) bool {
	return "" != r.request.Header.Get(k)
}

// GetHeader gets the first value associated with the given key.
func (r *Request) GetHeader(k string) string {
	return r.request.Header.Get(k)
}
