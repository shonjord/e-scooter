package middleware_test

import (
	"net/http"
	"strings"
	"testing"

	_http "github.com/shonjord/e-scooter/internal/pkg/presentation/http"
	"github.com/shonjord/e-scooter/internal/pkg/presentation/http/middleware"
	assert "github.com/shonjord/e-scooter/test/.assert"
	mock "github.com/shonjord/e-scooter/test/.mock"
)

// TestXApiKeyAuthenticator responsible to test different important scenarios of this middleware.
func TestXApiKeyAuthenticator(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(*testing.T, string)
	}{
		{
			// Given a valid request and response
			scenario: "When request doesn't contain header x-api-key",
			function: thenAuthenticatorMiddlewareShouldReturnErrorWhenXAPIKeyIsMissingFromRequest,
		},
		{
			// Given a valid request and response
			scenario: "When request contains an x-api-key that is not recognized by the application",
			function: thenAuthenticatorMiddlewareShouldReturnErrorWhenXAPIKeyIsNotRecognized,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(t, "xAPIKey")
		})
	}
}

func thenAuthenticatorMiddlewareShouldReturnErrorWhenXAPIKeyIsMissingFromRequest(t *testing.T, xAPIKey string) {
	req, _ := http.NewRequest("post", "url", strings.NewReader("body"))

	m := middleware.NewMobileXApiKeyAuthenticator(xAPIKey)

	err := m.Handle(
		_http.NewRequest(req),
		_http.NewResponse(new(mock.ResponseWriter)),
	)

	assert.Equals(t, err, &_http.Error{
		HTTPStatusCounterpart: http.StatusUnauthorized,
		Message:               "action is not authorized",
	})
}

func thenAuthenticatorMiddlewareShouldReturnErrorWhenXAPIKeyIsNotRecognized(t *testing.T, xAPIKey string) {
	req, _ := http.NewRequest("post", "url", strings.NewReader("body"))
	req.Header.Set("x-api-key", "invalid-x-api-key")

	m := middleware.NewScooterXApiKeyAuthenticator(xAPIKey)

	err := m.Handle(
		_http.NewRequest(req),
		_http.NewResponse(new(mock.ResponseWriter)),
	)

	assert.Equals(t, err, &_http.Error{
		HTTPStatusCounterpart: http.StatusUnauthorized,
		Message:               "action is not authorized",
	})
}
