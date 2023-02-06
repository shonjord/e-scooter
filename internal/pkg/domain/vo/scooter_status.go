package vo

const (
	ScooterAvailable = "available"
	ScooterOccupied  = "occupied"
)

type (
	ScooterStatus string
)

// NewScooterStatusAvailable returns a new available status.
func NewScooterStatusAvailable() *ScooterStatus {
	status := ScooterStatus(ScooterAvailable)

	return &status
}

// NewScooterOccupiedStatus returns a new available status.
func NewScooterOccupiedStatus() *ScooterStatus {
	status := ScooterStatus(ScooterOccupied)

	return &status
}

// IsOccupied verifies if this status is occupied
func (s *ScooterStatus) IsOccupied() bool {
	return *s == ScooterOccupied
}
