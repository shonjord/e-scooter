package http

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

// Bridge is responsible to receive an HTTP handler and execute it before endpoint calls.
func Bridge(b httpHandler) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			req := NewRequest(r)
			res := NewResponse(w)

			if err := b.Handle(req, res); err != nil {
				httpError := toHttpError(err)

				res.SetJSONContentType()
				res.WriteHeader(httpError.HTTPStatusCounterpart)

				if err = res.WriteHttpError(httpError); err != nil {
					log.WithError(err).Error("while writing HTTP error")
				}

				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
