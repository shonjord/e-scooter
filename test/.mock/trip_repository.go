package mock

import (
	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
)

type (
	TripRepository struct {
		FindByUUIDFunc           func(uuid.UUID) (*entity.Trip, error)
		FindTripInProgressByFunc func(scooterUUID, mobileUUID uuid.UUID) (*entity.Trip, error)
		UpdateFunc               func(*entity.Trip) error
		SaveFunc                 func(*entity.Trip) error
	}
)

// FindByUUID refer to the consumer of the interface for documentation.
func (m *TripRepository) FindByUUID(tripUUID uuid.UUID) (*entity.Trip, error) {
	return m.FindByUUIDFunc(tripUUID)
}

// FindTripInProgressBy refer to the consumer of the interface for documentation.
func (m *TripRepository) FindTripInProgressBy(scooterUUID, mobileUUID uuid.UUID) (*entity.Trip, error) {
	return m.FindTripInProgressByFunc(scooterUUID, mobileUUID)
}

// Update refer to the consumer of the interface for documentation.
func (m *TripRepository) Update(t *entity.Trip) error {
	return m.UpdateFunc(t)
}

// Save refer to the consumer of the interface for documentation.
func (m *TripRepository) Save(t *entity.Trip) error {
	return m.SaveFunc(t)
}
