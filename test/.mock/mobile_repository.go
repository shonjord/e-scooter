package mock

import (
	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
)

type (
	MobileRepository struct {
		FindMobileByUUIDFunc func(uuid.UUID) (*entity.Mobile, error)
	}
)

// FindMobileByUUID refer to the consumer of the interface for documentation.
func (m *MobileRepository) FindMobileByUUID(mobileUUID uuid.UUID) (*entity.Mobile, error) {
	return m.FindMobileByUUIDFunc(mobileUUID)
}
