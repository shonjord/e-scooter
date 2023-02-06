package vo

const (
	TripInProgress = "in_progress"
	TripFinished   = "finished"
)

type (
	TripStatus string
)

// NewInProgressStatus returns a new trip status that is in progress.
func NewInProgressStatus() *TripStatus {
	status := TripStatus(TripInProgress)

	return &status
}

// NewFinishedStatus returns a new trip status that is finished.
func NewFinishedStatus() *TripStatus {
	status := TripStatus(TripFinished)

	return &status
}

// IsFinished verifies if the value of this status is in finished.
func (t *TripStatus) IsFinished() bool {
	return *t == TripFinished
}
