package mock

import (
	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
)

type (
	ScooterRepository struct {
		UpdateFunc                func(*entity.Scooter) error
		FindByUUIDFunc            func(uuid.UUID) (*entity.Scooter, error)
		FindAvailableScootersFunc func() ([]*entity.Scooter, error)
	}
)

// Update refer to the consumer of the interface for documentation.
func (m *ScooterRepository) Update(s *entity.Scooter) error {
	return m.UpdateFunc(s)
}

// FindByUUID refer to the consumer of the interface for documentation.
func (m *ScooterRepository) FindByUUID(scooterUUID uuid.UUID) (*entity.Scooter, error) {
	return m.FindByUUIDFunc(scooterUUID)
}

// FindAvailableScooters refer to the consumer of the interface for documentation.
func (m *ScooterRepository) FindAvailableScooters() ([]*entity.Scooter, error) {
	return m.FindAvailableScootersFunc()
}
