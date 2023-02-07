package errors

import (
	"fmt"

	"github.com/google/uuid"
)

// NewTripForUUIDNotFound returns an error for when a trip for the given UUID is not found.
func NewTripForUUIDNotFound(uuid uuid.UUID) *NotFoundError {
	return &NotFoundError{
		Message: fmt.Sprintf("trip with UUID: %s, not found", uuid),
	}
}
