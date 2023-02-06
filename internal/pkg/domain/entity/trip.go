package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/vo"
)

type (
	Trip struct {
		ID        int            `json:"-"`
		UUID      uuid.UUID      `json:"uuid"`
		Mobile    *Mobile        `json:"mobile"`
		Scooter   *Scooter       `json:"scooter"`
		Status    *vo.TripStatus `json:"status"`
		CreatedAt time.Time      `json:"created_at"`
		UpdatedAt time.Time      `json:"updated_at"`
	}
)

// NewTrip returns a new instance of a new trip.
func NewTrip(m *Mobile, s *Scooter) *Trip {
	return &Trip{
		UUID:      uuid.New(),
		Mobile:    m,
		Scooter:   s,
		Status:    vo.NewInProgressStatus(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

// UpdateStatusToFinished updates the state of this trip status to finished.
func (e *Trip) UpdateStatusToFinished() {
	e.Status = vo.NewFinishedStatus()
}

// HasFinished verifies if this trip has finished.
func (e *Trip) HasFinished() bool {
	return e.Status.IsFinished()
}
