package handler_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/application/handler"
	"github.com/shonjord/e-scooter/internal/pkg/domain/command"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
	assert "github.com/shonjord/e-scooter/test/.assert"
	mock "github.com/shonjord/e-scooter/test/.mock"
)

// TestDisconnectScooter responsible to test different important scenarios of this application handler.
func TestDisconnectScooter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(
			*testing.T,
			*mock.ScooterRepository,
			*mock.TripRepository,
			*command.DisconnectScooter,
		)
	}{
		{
			// Given a valid command
			scenario: "When trip repository can't find a trip for the given scooter and mobile UUIDs",
			function: thenDisconnectScooterShouldFailWithTripRepositoryError,
		},
		{
			// Given a valid command
			scenario: "When scooter repository can't update a scooter",
			function: thenDisconnectScooterShouldFailWithScooterRepositoryError,
		},
		{
			// Given a valid command
			scenario: "When trip repository can't update a trip",
			function: thenDisconnectScooterShouldFailWithTripRepositoryUpdateError,
		},
		{
			// Given a valid command
			scenario: "When disconnect scooter handler handles the command successfully",
			function: thenDisconnectScooterHandlerShouldReturnNoError,
		},
	}

	cmd := &command.DisconnectScooter{
		MobileUUID:  uuid.New(),
		ScooterUUID: uuid.New(),
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(
				t,
				new(mock.ScooterRepository),
				new(mock.TripRepository),
				cmd,
			)
		})
	}
}

func thenDisconnectScooterShouldFailWithTripRepositoryError(
	t *testing.T,
	sr *mock.ScooterRepository,
	tr *mock.TripRepository,
	cmd *command.DisconnectScooter,
) {
	tr.FindTripInProgressByFunc = func(scooterUUID, mobileUUID uuid.UUID) (*entity.Trip, error) {
		return nil, errors.New("trip not found")
	}

	h := handler.NewDisconnectScooter(tr, sr)
	err := h.DisconnectScooter(cmd)

	assert.True(t, err != nil, "trip repository should return errors when trip is not found.")
}

func thenDisconnectScooterShouldFailWithScooterRepositoryError(
	t *testing.T,
	sr *mock.ScooterRepository,
	tr *mock.TripRepository,
	cmd *command.DisconnectScooter,
) {
	tr.FindTripInProgressByFunc = func(scooterUUID, mobileUUID uuid.UUID) (*entity.Trip, error) {
		trip := new(entity.Trip)
		trip.Scooter = new(entity.Scooter)
		trip.UUID = cmd.MobileUUID

		return trip, nil
	}

	sr.UpdateFunc = func(*entity.Scooter) error {
		return errors.New("errors while updating scooter")
	}

	h := handler.NewDisconnectScooter(tr, sr)
	err := h.DisconnectScooter(cmd)

	assert.True(t, err != nil, "scooter repository should return errors when something unexpected happens.")
}

func thenDisconnectScooterShouldFailWithTripRepositoryUpdateError(
	t *testing.T,
	sr *mock.ScooterRepository,
	tr *mock.TripRepository,
	cmd *command.DisconnectScooter,
) {
	tr.FindTripInProgressByFunc = func(scooterUUID, mobileUUID uuid.UUID) (*entity.Trip, error) {
		trip := new(entity.Trip)
		trip.Scooter = new(entity.Scooter)
		trip.UUID = cmd.MobileUUID

		return trip, nil
	}

	sr.UpdateFunc = func(*entity.Scooter) error {
		return nil
	}

	tr.UpdateFunc = func(*entity.Trip) error {
		return errors.New("errors while updating a trip")
	}

	h := handler.NewDisconnectScooter(tr, sr)
	err := h.DisconnectScooter(cmd)

	assert.True(t, err != nil, "trip repository should return errors when something unexpected happens.")
}

func thenDisconnectScooterHandlerShouldReturnNoError(
	t *testing.T,
	sr *mock.ScooterRepository,
	tr *mock.TripRepository,
	cmd *command.DisconnectScooter,
) {
	tr.FindTripInProgressByFunc = func(scooterUUID, mobileUUID uuid.UUID) (*entity.Trip, error) {
		trip := new(entity.Trip)
		trip.Scooter = new(entity.Scooter)
		trip.UUID = cmd.MobileUUID

		return trip, nil
	}

	sr.UpdateFunc = func(*entity.Scooter) error {
		return nil
	}

	tr.UpdateFunc = func(*entity.Trip) error {
		return nil
	}

	h := handler.NewDisconnectScooter(tr, sr)
	err := h.DisconnectScooter(cmd)

	assert.True(t, err == nil, "handler should handle command successfully for this scenario.")
}
