package handler

import "github.com/shonjord/e-scooter/internal/pkg/domain/entity"

type (
	scootersFinder interface {
		FindAvailableScooters() ([]*entity.Scooter, error)
	}

	FindAvailableScooters struct {
		finder scootersFinder
	}
)

// NewFindAvailableScooters finds all possible scooters available to ride.
func NewFindAvailableScooters(f scootersFinder) *FindAvailableScooters {
	return &FindAvailableScooters{
		finder: f,
	}
}

// FindAvailableScooters attempts to find all available scooters to ride.
func (h *FindAvailableScooters) FindAvailableScooters() ([]*entity.Scooter, error) {
	return h.finder.FindAvailableScooters()
}
