package mysql

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
)

type (
	MobileRepository struct {
		connection *Connection
	}
)

// NewMobileRepository returns a new instance of this repository.
func NewMobileRepository(c *Connection) *MobileRepository {
	return &MobileRepository{
		connection: c,
	}
}

// FindMobileByUUID finds a mobile client for the given UUID.
func (m *MobileRepository) FindMobileByUUID(uuid uuid.UUID) (*entity.Mobile, error) {
	var (
		mobile = new(entity.Mobile)
		query  = fmt.Sprintf("SELECT * FROM mobiles WHERE uuid = '%s'", uuid)
	)

	if err := m.connection.FindOneBy(query, mobileDBValues(mobile)...); err != nil {
		return nil, err
	}

	return mobile, nil
}

// mobileDBValues returns all DB values of the mobile entity.
func mobileDBValues(m *entity.Mobile) []interface{} {
	return []interface{}{
		&m.ID,
		&m.UUID,
		&m.CreatedAt,
		&m.UpdatedAt,
	}
}
