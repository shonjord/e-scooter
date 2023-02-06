package handler

import (
	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/domain/command"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
	"github.com/shonjord/e-scooter/internal/pkg/domain/errors"
	"github.com/shonjord/e-scooter/internal/pkg/domain/event"
)

type (
	scooterFinderUpdater interface {
		FindByUUID(uuid.UUID) (*entity.Scooter, error)
		Update(*entity.Scooter) error
	}
	mobileFinder interface {
		FindMobileByUUID(uuid.UUID) (*entity.Mobile, error)
	}
	tripSaver interface {
		Save(*entity.Trip) error
	}
	eventDispatcher interface {
		Dispatch(*event.DomainEvent) error
	}

	ConnectScooter struct {
		scooterFinderUpdater scooterFinderUpdater
		mobileFinder         mobileFinder
		tripSaver            tripSaver
		dispatcher           eventDispatcher
	}
)

// NewConnectScooter returns a new instance of this scooter connector.
func NewConnectScooter(
	sfu scooterFinderUpdater,
	mf mobileFinder,
	ts tripSaver,
	d eventDispatcher,
) *ConnectScooter {
	return &ConnectScooter{
		scooterFinderUpdater: sfu,
		mobileFinder:         mf,
		tripSaver:            ts,
		dispatcher:           d,
	}
}

// ConnectScooter connects a scooter with a mobile client for a new trip :).
func (h *ConnectScooter) ConnectScooter(c *command.ConnectScooter) error {
	mobile, err := h.mobileFinder.FindMobileByUUID(c.MobileUUID)
	if err != nil {
		return err
	}

	scooter, err := h.scooterFinderUpdater.FindByUUID(c.ScooterUUID)
	if err != nil {
		return err
	}

	if scooter.IsOccupied() {
		return errors.NewScooterIsOccupiedError(scooter)
	}

	scooter.UpdateStatusToOccupied()

	if err = h.scooterFinderUpdater.Update(scooter); err != nil {
		return err
	}

	trip := entity.NewTrip(mobile, scooter)

	if err = h.tripSaver.Save(trip); err != nil {
		return err
	}

	if err = h.dispatcher.Dispatch(event.NewTripStarted(trip)); err != nil {
		return err
	}

	return nil
}
