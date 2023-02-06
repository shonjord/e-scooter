package mock

import "github.com/shonjord/e-scooter/internal/pkg/domain/vo"

type (
	GeographicLocator struct {
		FindCurrentGeographicLocationFunc func() (*vo.Coordinate, *vo.Coordinate, error)
	}
)

// FindCurrentGeographicLocation refer to the consumer of the interface for documentation.
func (m *GeographicLocator) FindCurrentGeographicLocation() (*vo.Coordinate, *vo.Coordinate, error) {
	return m.FindCurrentGeographicLocationFunc()
}
