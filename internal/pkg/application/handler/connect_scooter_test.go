package handler_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/application/handler"
	"github.com/shonjord/e-scooter/internal/pkg/domain/command"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
	"github.com/shonjord/e-scooter/internal/pkg/domain/event"
	"github.com/shonjord/e-scooter/internal/pkg/domain/vo"
	assert "github.com/shonjord/e-scooter/test/.assert"
	mock "github.com/shonjord/e-scooter/test/.mock"
)

// TestConnectScooter responsible to test different important scenarios of this application handler.
func TestConnectScooter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(
			*testing.T,
			*mock.ScooterRepository,
			*mock.MobileRepository,
			*mock.TripRepository,
			*mock.EventDispatcher,
			*command.ConnectScooter,
		)
	}{
		{
			// Given a valid command
			scenario: "When mobile repository can't find a mobile",
			function: thenConnectScooterShouldReturnMobileNotFoundError,
		},
		{
			// Given a valid command
			scenario: "When scooter repository can't find a scooter",
			function: thenConnectScooterShouldReturnScooterNotFoundError,
		},
		{
			// Given a valid command
			scenario: "When scooter repository find a scooter that is occupied",
			function: thenHandlerShouldReturnScooterIsOccupiedError,
		},
		{
			// Given a valid command
			scenario: "When scooter repository can't update a scooter",
			function: thenConnectScooterShouldReturnScooterRepositoryError,
		},
		{
			// Given a valid command
			scenario: "When trip repository can't save a trip",
			function: thenConnectScooterShouldReturnTripRepositoryError,
		},
		{
			// Given a valid command
			scenario: "When dispatcher can't dispatch a message",
			function: thenConnectScooterShouldReturnDispatcherError,
		},
		{
			// Given a valid command
			scenario: "When handler handles command successfully",
			function: thenConnectScooterHandlerShouldReturnNoError,
		},
	}

	cmd := &command.ConnectScooter{
		MobileUUID:  uuid.New(),
		ScooterUUID: uuid.New(),
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(
				t,
				new(mock.ScooterRepository),
				new(mock.MobileRepository),
				new(mock.TripRepository),
				new(mock.EventDispatcher),
				cmd,
			)
		})
	}
}

func thenConnectScooterShouldReturnMobileNotFoundError(
	t *testing.T,
	sr *mock.ScooterRepository,
	mr *mock.MobileRepository,
	tr *mock.TripRepository,
	ed *mock.EventDispatcher,
	cmd *command.ConnectScooter,
) {
	mr.FindMobileByUUIDFunc = func(u uuid.UUID) (*entity.Mobile, error) {
		return nil, errors.New("mobile not found")
	}

	h := handler.NewConnectScooter(sr, mr, tr, ed)
	err := h.ConnectScooter(cmd)

	assert.True(t, err != nil, "should fail with mobile repository errors.")
}

func thenConnectScooterShouldReturnScooterNotFoundError(
	t *testing.T,
	sr *mock.ScooterRepository,
	mr *mock.MobileRepository,
	tr *mock.TripRepository,
	ed *mock.EventDispatcher,
	cmd *command.ConnectScooter,
) {
	mr.FindMobileByUUIDFunc = func(uuid.UUID) (*entity.Mobile, error) {
		return new(entity.Mobile), nil
	}

	sr.FindByUUIDFunc = func(uuid.UUID) (*entity.Scooter, error) {
		return nil, errors.New("scooter not found")
	}

	h := handler.NewConnectScooter(sr, mr, tr, ed)
	err := h.ConnectScooter(cmd)

	assert.True(t, err != nil, "should fail with scooter repository errors.")
}

func thenHandlerShouldReturnScooterIsOccupiedError(
	t *testing.T,
	sr *mock.ScooterRepository,
	mr *mock.MobileRepository,
	tr *mock.TripRepository,
	ed *mock.EventDispatcher,
	cmd *command.ConnectScooter,
) {
	mr.FindMobileByUUIDFunc = func(uuid.UUID) (*entity.Mobile, error) {
		return new(entity.Mobile), nil
	}

	sr.FindByUUIDFunc = func(uuid.UUID) (*entity.Scooter, error) {
		scooter := new(entity.Scooter)
		scooter.Status = vo.NewScooterOccupiedStatus()

		return scooter, nil
	}

	h := handler.NewConnectScooter(sr, mr, tr, ed)
	err := h.ConnectScooter(cmd)

	assert.True(t, err != nil, "should fail with scooter is occupied.")
}

func thenConnectScooterShouldReturnScooterRepositoryError(
	t *testing.T,
	sr *mock.ScooterRepository,
	mr *mock.MobileRepository,
	tr *mock.TripRepository,
	ed *mock.EventDispatcher,
	cmd *command.ConnectScooter,
) {
	mr.FindMobileByUUIDFunc = func(uuid.UUID) (*entity.Mobile, error) {
		return new(entity.Mobile), nil
	}

	sr.FindByUUIDFunc = func(uuid.UUID) (*entity.Scooter, error) {
		scooter := new(entity.Scooter)
		scooter.Status = vo.NewScooterStatusAvailable()

		return scooter, nil
	}

	sr.UpdateFunc = func(s *entity.Scooter) error {
		return errors.New("internal mobile")
	}

	h := handler.NewConnectScooter(sr, mr, tr, ed)
	err := h.ConnectScooter(cmd)

	assert.True(t, err != nil, "should fail with scooter repository update errors.")
}

func thenConnectScooterShouldReturnTripRepositoryError(
	t *testing.T,
	sr *mock.ScooterRepository,
	mr *mock.MobileRepository,
	tr *mock.TripRepository,
	ed *mock.EventDispatcher,
	cmd *command.ConnectScooter,
) {
	mr.FindMobileByUUIDFunc = func(uuid.UUID) (*entity.Mobile, error) {
		return new(entity.Mobile), nil
	}

	sr.FindByUUIDFunc = func(uuid.UUID) (*entity.Scooter, error) {
		scooter := new(entity.Scooter)
		scooter.Status = vo.NewScooterStatusAvailable()

		return scooter, nil
	}

	sr.UpdateFunc = func(s *entity.Scooter) error {
		return nil
	}

	tr.SaveFunc = func(*entity.Trip) error {
		return errors.New("mobile while saving a trip")
	}

	h := handler.NewConnectScooter(sr, mr, tr, ed)
	err := h.ConnectScooter(cmd)

	assert.True(t, err != nil, "should fail with trip repository save errors.")
}

func thenConnectScooterShouldReturnDispatcherError(
	t *testing.T,
	sr *mock.ScooterRepository,
	mr *mock.MobileRepository,
	tr *mock.TripRepository,
	ed *mock.EventDispatcher,
	cmd *command.ConnectScooter,
) {
	mr.FindMobileByUUIDFunc = func(uuid.UUID) (*entity.Mobile, error) {
		return new(entity.Mobile), nil
	}

	sr.FindByUUIDFunc = func(uuid.UUID) (*entity.Scooter, error) {
		scooter := new(entity.Scooter)
		scooter.Status = vo.NewScooterStatusAvailable()

		return scooter, nil
	}

	sr.UpdateFunc = func(s *entity.Scooter) error {
		return nil
	}

	tr.SaveFunc = func(*entity.Trip) error {
		return nil
	}

	ed.DispatchFunc = func(*event.DomainEvent) error {
		return errors.New("mobile dispatching message")
	}

	h := handler.NewConnectScooter(sr, mr, tr, ed)
	err := h.ConnectScooter(cmd)

	assert.True(t, err != nil, "should fail with event dispatcher errors.")
}

func thenConnectScooterHandlerShouldReturnNoError(
	t *testing.T,
	sr *mock.ScooterRepository,
	mr *mock.MobileRepository,
	tr *mock.TripRepository,
	ed *mock.EventDispatcher,
	cmd *command.ConnectScooter,
) {
	mr.FindMobileByUUIDFunc = func(uuid.UUID) (*entity.Mobile, error) {
		return new(entity.Mobile), nil
	}

	sr.FindByUUIDFunc = func(uuid.UUID) (*entity.Scooter, error) {
		scooter := new(entity.Scooter)
		scooter.Status = vo.NewScooterStatusAvailable()

		return scooter, nil
	}

	sr.UpdateFunc = func(s *entity.Scooter) error {
		return nil
	}

	tr.SaveFunc = func(*entity.Trip) error {
		return nil
	}

	ed.DispatchFunc = func(*event.DomainEvent) error {
		return nil
	}

	h := handler.NewConnectScooter(sr, mr, tr, ed)
	err := h.ConnectScooter(cmd)

	assert.True(t, err == nil, "connect scooter handler should not fail.")
}
