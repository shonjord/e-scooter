package event

import (
	"time"

	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
	"github.com/shonjord/e-scooter/internal/pkg/domain/event"
	"github.com/shonjord/e-scooter/internal/pkg/domain/vo"
	log "github.com/sirupsen/logrus"
)

type (
	scooterUpdater interface {
		Update(*entity.Scooter) error
	}
	tripFinderUpdater interface {
		FindByUUID(uuid.UUID) (*entity.Trip, error)
		Update(t *entity.Trip) error
	}
	geographicLocator interface {
		FindCurrentGeographicLocation() (*vo.Coordinate, *vo.Coordinate, error)
	}

	TripStartedEventHandler struct {
		finderUpdater tripFinderUpdater
		updater       scooterUpdater
		locator       geographicLocator
	}
)

// NewTripStartedEventHandler returns an ew instance of this event handler.
func NewTripStartedEventHandler(
	fu tripFinderUpdater,
	u scooterUpdater,
	l geographicLocator,
) *TripStartedEventHandler {
	return &TripStartedEventHandler{
		finderUpdater: fu,
		updater:       u,
		locator:       l,
	}
}

// SubscribedTo verifies to whom is this event handler subscribed to.
func (e *TripStartedEventHandler) SubscribedTo() string {
	return event.TripStarted
}

// Handle verifies the status of the trip, if the trip finished, then the event finishes.
// if the trip is in progress, an update to scooter with its latest geographic location is applied.
func (e *TripStartedEventHandler) Handle(event *event.DomainEvent) error {
	for {
		trip, err := e.findTrip(event.UUID)
		if err != nil {
			return err
		}

		e.updateScooterCoordinates(trip.Scooter)

		if trip.HasFinished() {
			log.Infof("trip with UUID: %s, finished", trip.UUID)

			break
		}

		time.Sleep(5 * time.Second)
	}

	return nil
}

// findTrip finds a trip for the given UUID.
func (e *TripStartedEventHandler) findTrip(tripUUID uuid.UUID) (*entity.Trip, error) {
	trip, err := e.finderUpdater.FindByUUID(tripUUID)
	if err != nil {
		log.WithError(err).Errorf("errors while fetching trip with UUID: %s.", tripUUID)

		return nil, err
	}

	return trip, nil
}

// updateScooterCoordinates finds the current geographical location
// and applies it to the given scooter for updates.
func (e *TripStartedEventHandler) updateScooterCoordinates(s *entity.Scooter) {
	latitude, longitude, err := e.locator.FindCurrentGeographicLocation()
	if err != nil {
		log.WithError(err).Warnf(
			"could not find current geographic location to update scooter with UUID: %s.",
			s.UUID,
		)

		return
	}

	s.UpdateCoordinates(latitude, longitude)

	if err = e.updater.Update(s); err != nil {
		log.WithError(err).Warnf(
			"could not update current geographic location of scooter with UUID: %s.",
			s.UUID,
		)

		return
	}

	log.Infof(
		"scooter with uuid: %s, is currently located at: latitude: %d, longitude: %d.",
		s.UUID,
		s.Latitude,
		s.Longitude,
	)
}
