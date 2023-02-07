package mysql

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
	"github.com/shonjord/e-scooter/internal/pkg/domain/errors"
)

type (
	ScooterRepository struct {
		connection *Connection
	}
)

// NewScooterRepository returns a new instance of this repository.
func NewScooterRepository(c *Connection) *ScooterRepository {
	return &ScooterRepository{
		connection: c,
	}
}

// FindAvailableScooters queries the available scooters of the service.
func (m *ScooterRepository) FindAvailableScooters() ([]*entity.Scooter, error) {
	var (
		scooters []*entity.Scooter
		query    = "SELECT * FROM scooters WHERE status = 'available'"
	)

	callback := func(rows *sql.Rows) error {
		scooter := new(entity.Scooter)

		if err := rows.Scan(scooterDBValues(scooter)...); err != nil {
			return err
		}

		scooters = append(scooters, scooter)

		return nil
	}

	if err := m.connection.FindManyBy(query, callback); err != nil {
		return nil, err
	}

	return scooters, nil
}

// FindByUUID finds a scooter for the given UUID.
func (m *ScooterRepository) FindByUUID(uuid uuid.UUID) (*entity.Scooter, error) {
	var (
		scooter = new(entity.Scooter)
		query   = fmt.Sprintf("SELECT * FROM scooters WHERE uuid = '%s'", uuid)
	)

	if err := m.connection.FindOneBy(query, scooterDBValues(scooter)...); err != nil {
		if errorIsNoRows(err) {
			return nil, errors.NewScooterForUUIDNotFound(uuid)
		}
		return nil, err
	}

	return scooter, nil
}

// Update executes an UPDATE statement for the given scooter.
func (m *ScooterRepository) Update(s *entity.Scooter) error {
	var (
		query = "UPDATE scooters SET latitude = ?, longitude = ?, status = ?, updated_at = ? WHERE id = ?"
	)

	return m.connection.Execute(query, s.Latitude, s.Longitude, s.Status, time.Now(), s.ID)
}

// scooterDBValues returns all DB values of the scooter entity.
func scooterDBValues(s *entity.Scooter) []interface{} {
	return []interface{}{
		&s.ID,
		&s.UUID,
		&s.Latitude,
		&s.Longitude,
		&s.Status,
		&s.CreatedAt,
		&s.UpdatedAt,
	}
}
