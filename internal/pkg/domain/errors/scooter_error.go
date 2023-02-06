package errors

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
)

// NewScooterIsOccupiedError returns an errors for when a scooter can't be taken because is occupied.
func NewScooterIsOccupiedError(s *entity.Scooter) *ConflictError {
	return &ConflictError{
		Message: fmt.Sprintf("the following scooter is occupied: %s", s.UUID),
	}
}

// NewScooterForUUIDNotFound returns an error for when a scooter for the given UUID is not found.
func NewScooterForUUIDNotFound(uuid uuid.UUID) *NotFoundError {
	return &NotFoundError{
		Message: fmt.Sprintf("scooter with UUID: %s, not found", uuid),
	}
}
