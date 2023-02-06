package handler

import (
	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/command"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
	log "github.com/sirupsen/logrus"
)

type (
	scooterUpdater interface {
		Update(*entity.Scooter) error
	}
	tripFinderUpdater interface {
		FindTripInProgressBy(scooterUUID, mobileUUID uuid.UUID) (*entity.Trip, error)
		Update(trip *entity.Trip) error
	}

	DisconnectScooter struct {
		finderUpdater tripFinderUpdater
		updater       scooterUpdater
	}
)

// NewDisconnectScooter returns a new instance of a scooter disconnection application handler.
func NewDisconnectScooter(fu tripFinderUpdater, u scooterUpdater) *DisconnectScooter {
	return &DisconnectScooter{
		finderUpdater: fu,
		updater:       u,
	}
}

// DisconnectScooter disconnect mobile from scooter, making scooter available for pickup :).
func (h *DisconnectScooter) DisconnectScooter(c *command.DisconnectScooter) error {
	trip, err := h.finderUpdater.FindTripInProgressBy(c.ScooterUUID, c.MobileUUID)
	if err != nil {
		return err
	}

	log.Infof("updating scooter with uuid: %s, status to available", trip.UUID)
	scooter := trip.Scooter
	scooter.UpdateStatusToAvailable()

	if err = h.updater.Update(scooter); err != nil {
		return err
	}

	log.Infof("updating trip with uuid: %s, to finished", scooter.UUID)
	trip.UpdateStatusToFinished()

	if err = h.finderUpdater.Update(trip); err != nil {
		return err
	}

	return nil
}
