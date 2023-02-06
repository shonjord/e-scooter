package mock

import "net/http"

type (
	// ResponseWriter is a mock of an actual write of HTTP responses.
	ResponseWriter struct {
		HeaderFunc      func() http.Header
		WriteFunc       func([]byte) (int, error)
		WriteHeaderFunc func(statusCode int)
	}
)

// Header refer to the consumer of the interface for documentation.
func (r *ResponseWriter) Header() http.Header {
	return r.HeaderFunc()
}

// Write refer to the consumer of the interface for documentation.
func (r *ResponseWriter) Write(c []byte) (int, error) {
	return r.WriteFunc(c)
}

// WriteHeader refer to the consumer of the interface for documentation.
func (r *ResponseWriter) WriteHeader(statusCode int) {
	r.WriteHeaderFunc(statusCode)
}
