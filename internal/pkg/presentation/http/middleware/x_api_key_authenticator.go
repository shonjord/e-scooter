package middleware

import (
	"github.com/shonjord/e-scooter/internal/pkg/presentation/http"
)

const (
	headerKey = "x-api-key"
)

type (
	XApiKeyAuthenticator struct {
		xApiKey string
	}
)

// NewMobileXApiKeyAuthenticator returns a new middleware to authenticate mobile endpoints.
func NewMobileXApiKeyAuthenticator(mobileXApiKey string) *XApiKeyAuthenticator {
	return &XApiKeyAuthenticator{
		xApiKey: mobileXApiKey,
	}
}

// NewScooterXApiKeyAuthenticator returns a new middleware to authenticate scooter endpoints.
func NewScooterXApiKeyAuthenticator(scooterXApiKey string) *XApiKeyAuthenticator {
	return &XApiKeyAuthenticator{
		xApiKey: scooterXApiKey,
	}
}

// Handle verifies if the given request is authorized.
func (m *XApiKeyAuthenticator) Handle(req *http.Request, res *http.Response) error {
	if !req.HasHeader(headerKey) {
		return res.UnauthorizedRequest()
	}

	if req.GetHeader(headerKey) != m.xApiKey {
		return res.UnauthorizedRequest()
	}

	return nil
}
