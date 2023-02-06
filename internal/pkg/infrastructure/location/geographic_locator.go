package location

import (
	"math/rand"

	"github.com/shonjord/e-scooter/internal/pkg/domain/vo"
)

type (
	GeographicLocator struct{}
)

// NewGeographicLocator is responsible to find geographic location.
func NewGeographicLocator() *GeographicLocator {
	return new(GeographicLocator)
}

// FindCurrentGeographicLocation fetches the current geographic location.
// values are mocked for prototyping purposes.
func (l *GeographicLocator) FindCurrentGeographicLocation() (*vo.Coordinate, *vo.Coordinate, error) {
	return vo.NewCoordinate(rand.Float64()), vo.NewCoordinate(rand.Float64()), nil
}
