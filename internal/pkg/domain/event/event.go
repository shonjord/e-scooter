package event

import (
	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
)

const (
	TripStarted = "TripStarted"
)

type (
	DomainEvent struct {
		UUID  uuid.UUID
		Name  string
		Async bool
	}
)

// NewTripStarted returns a new event when a trip has been started (to be handled async).
func NewTripStarted(t *entity.Trip) *DomainEvent {
	return &DomainEvent{
		UUID:  t.UUID,
		Name:  TripStarted,
		Async: true,
	}
}
