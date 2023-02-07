package vo

type (
	Coordinate float64
)

// NewCoordinate returns a new coordinate with the given value.
func NewCoordinate(value float64) *Coordinate {
	coordinate := Coordinate(value)

	return &coordinate
}

// ToFloat returns a float(base64) representation of this coordinate.
func (c *Coordinate) ToFloat() float64 {
	return float64(*c)
}
