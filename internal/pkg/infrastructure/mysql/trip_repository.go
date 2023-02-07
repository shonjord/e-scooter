package mysql

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
	"github.com/shonjord/e-scooter/internal/pkg/domain/errors"
)

type (
	TripRepository struct {
		connection *Connection
	}
)

// NewTripRepository returns a new trip repository, responsible to handle all queries related to a trip.
func NewTripRepository(c *Connection) *TripRepository {
	return &TripRepository{
		connection: c,
	}
}

// FindByUUID finds a trip by the given UUID.
func (m *TripRepository) FindByUUID(tripUUID uuid.UUID) (*entity.Trip, error) {
	var (
		trip    = new(entity.Trip)
		scooter = new(entity.Scooter)
		mobile  = new(entity.Mobile)
		query   = fmt.Sprintf("SELECT * FROM trips WHERE uuid = '%s'", tripUUID)
	)

	if err := m.connection.FindOneBy(query, tripDBValues(trip, scooter, mobile)...); err != nil {
		if errorIsNoRows(err) {
			return nil, errors.NewTripForUUIDNotFound(tripUUID)
		}
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM scooters WHERE uuid = '%s'", scooter.UUID)
	if err := m.connection.FindOneBy(query, scooterDBValues(scooter)...); err != nil {
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM mobiles WHERE uuid = '%s'", mobile.UUID)
	if err := m.connection.FindOneBy(query, mobileDBValues(mobile)...); err != nil {
		return nil, err
	}

	trip.Scooter = scooter
	trip.Mobile = mobile

	return trip, nil
}

// Save persist a new trip into MySQL DB.
func (m *TripRepository) Save(t *entity.Trip) error {
	query := "INSERT INTO trips (uuid, scooter_uuid, mobile_uuid, status, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	smt, err := m.connection.Prepare(query)
	if err != nil {
		return err
	}

	values := []interface{}{
		t.UUID,
		t.Scooter.UUID,
		t.Mobile.UUID,
		t.Status,
		t.CreatedAt,
		t.UpdatedAt,
	}

	_, err = smt.Exec(values...)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return nil
}

// Update executes an UPDATE statement for the given trip.
func (m *TripRepository) Update(t *entity.Trip) error {
	var (
		query = "UPDATE trips SET status = ?, updated_at = ? WHERE id = ?"
	)

	return m.connection.Execute(query, t.Status, time.Now(), t.ID)
}

// FindTripInProgressBy queries the persisted trips that are in progress for the given scooter and mobile UUIDs.
func (m *TripRepository) FindTripInProgressBy(scooterUUID, mobileUUID uuid.UUID) (*entity.Trip, error) {
	var (
		trip    = new(entity.Trip)
		scooter = new(entity.Scooter)
		mobile  = new(entity.Mobile)
		query   = fmt.Sprintf(
			"SELECT * FROM trips WHERE scooter_uuid = '%s' AND mobile_uuid = '%s' AND status = 'in_progress'",
			scooterUUID,
			mobileUUID,
		)
	)

	if err := m.connection.FindOneBy(query, tripDBValues(trip, scooter, mobile)...); err != nil {
		if errorIsNoRows(err) {
			return nil, &errors.NotFoundError{
				Message: fmt.Sprintf(
					"trip in progress for scooter UUID: %s, and mobile UUID: %s, not found",
					scooterUUID,
					mobileUUID,
				),
			}
		}
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM scooters WHERE uuid = '%s'", scooterUUID)
	if err := m.connection.FindOneBy(query, scooterDBValues(scooter)...); err != nil {
		return nil, err
	}

	query = fmt.Sprintf("SELECT * FROM mobiles WHERE uuid = '%s'", mobileUUID)
	if err := m.connection.FindOneBy(query, mobileDBValues(mobile)...); err != nil {
		return nil, err
	}

	trip.Scooter = scooter
	trip.Mobile = mobile

	return trip, nil

}

// tripDBValues returns all DB values of the trip entity.
func tripDBValues(t *entity.Trip, s *entity.Scooter, m *entity.Mobile) []interface{} {
	return []interface{}{
		&t.ID,
		&t.UUID,
		&s.UUID,
		&m.UUID,
		&t.Status,
		&t.CreatedAt,
		&t.UpdatedAt,
	}
}
