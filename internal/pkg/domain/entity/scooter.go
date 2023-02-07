package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/vo"
)

type (
	Scooter struct {
		ID        int               `json:"-"`
		UUID      uuid.UUID         `json:"uuid"`
		Latitude  *vo.Coordinate    `json:"latitude"`
		Longitude *vo.Coordinate    `json:"longitude"`
		Status    *vo.ScooterStatus `json:"status"`
		CreatedAt time.Time         `json:"-"`
		UpdatedAt time.Time         `json:"-"`
	}
)

// UpdateCoordinates updates this scooter with the given coordinates.
func (e *Scooter) UpdateCoordinates(la *vo.Coordinate, lo *vo.Coordinate) {
	e.Latitude = la
	e.Longitude = lo
}

// IsOccupied verifies if this scooter is occupied.
func (e *Scooter) IsOccupied() bool {
	return e.Status.IsOccupied()
}

// UpdateStatusToOccupied updates the state of this scooter status to occupied.
func (e *Scooter) UpdateStatusToOccupied() {
	e.Status = vo.NewScooterOccupiedStatus()
}

// UpdateStatusToAvailable updates the state of this scooter status to occupied.
func (e *Scooter) UpdateStatusToAvailable() {
	e.Status = vo.NewScooterStatusAvailable()
}
