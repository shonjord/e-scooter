package event_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/application/event"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
	domainEvent "github.com/shonjord/e-scooter/internal/pkg/domain/event"
	"github.com/shonjord/e-scooter/internal/pkg/domain/vo"
	assert "github.com/shonjord/e-scooter/test/.assert"
	mock "github.com/shonjord/e-scooter/test/.mock"
)

// TestTripStartedEventHandler responsible to test different important scenarios of this event handler.
func TestTripStartedEventHandler(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(
			*testing.T,
			*mock.TripRepository,
			*mock.ScooterRepository,
			*mock.GeographicLocator,
			*domainEvent.DomainEvent,
		)
	}{
		{
			// Given a valid event
			scenario: "When trip repository can't find a trip",
			function: thenTripEventHandlerShouldReturnTripRepositoryError,
		},
		{
			// Given a valid event
			scenario: "When geographic locator returns an errors",
			function: thenTripEventHandlerShouldNotUpdateScooterCoordinates,
		},
		{
			// Given a valid event
			scenario: "When geographic locator returns coordinates but scooter repository returns an errors",
			function: thenTripEventHandlerShouldNotUpdateScooterCoordinatesWhenScooterRepositoryReturnsAnError,
		},
		{
			// Given a valid event
			scenario: "When geographic locator returns coordinates",
			function: thenTripEventHandlerShouldNotReturnAnError,
		},
	}

	de := &domainEvent.DomainEvent{
		UUID:  uuid.New(),
		Name:  "event",
		Async: false,
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(t, new(mock.TripRepository), new(mock.ScooterRepository), new(mock.GeographicLocator), de)
		})
	}
}

func thenTripEventHandlerShouldReturnTripRepositoryError(
	t *testing.T,
	tr *mock.TripRepository,
	sr *mock.ScooterRepository,
	gl *mock.GeographicLocator,
	de *domainEvent.DomainEvent,
) {
	tr.FindByUUIDFunc = func(u uuid.UUID) (*entity.Trip, error) {
		return nil, errors.New("trip not found")
	}

	handler := event.NewTripStartedEventHandler(tr, sr, gl)
	err := handler.Handle(de)

	assert.True(t, err != nil, "when trip repository returns errors, handler should fail.")
}

func thenTripEventHandlerShouldNotUpdateScooterCoordinates(
	t *testing.T,
	tr *mock.TripRepository,
	sr *mock.ScooterRepository,
	gl *mock.GeographicLocator,
	de *domainEvent.DomainEvent,
) {
	latitude := vo.NewCoordinate(100)
	longitude := vo.NewCoordinate(200)

	scooter := new(entity.Scooter)
	scooter.UpdateCoordinates(latitude, longitude)

	trip := new(entity.Trip)
	trip.Scooter = scooter
	trip.Status = vo.NewFinishedStatus()

	tr.FindByUUIDFunc = func(u uuid.UUID) (*entity.Trip, error) {
		return trip, nil
	}

	gl.FindCurrentGeographicLocationFunc = func() (*vo.Coordinate, *vo.Coordinate, error) {
		return nil, nil, errors.New("errors while fetching current location")
	}

	handler := event.NewTripStartedEventHandler(tr, sr, gl)
	err := handler.Handle(de)
	if err != nil {
		t.Fail()

		return
	}

	assert.Equals(t, latitude, scooter.Latitude)
	assert.Equals(t, longitude, scooter.Longitude)
}

func thenTripEventHandlerShouldNotUpdateScooterCoordinatesWhenScooterRepositoryReturnsAnError(
	t *testing.T,
	tr *mock.TripRepository,
	sr *mock.ScooterRepository,
	gl *mock.GeographicLocator,
	de *domainEvent.DomainEvent,
) {
	latitude := vo.NewCoordinate(100)
	longitude := vo.NewCoordinate(200)

	scooter := new(entity.Scooter)
	scooter.UpdateCoordinates(latitude, longitude)

	trip := new(entity.Trip)
	trip.Scooter = scooter
	trip.Status = vo.NewFinishedStatus()

	tr.FindByUUIDFunc = func(u uuid.UUID) (*entity.Trip, error) {
		return trip, nil
	}

	gl.FindCurrentGeographicLocationFunc = func() (*vo.Coordinate, *vo.Coordinate, error) {
		return vo.NewCoordinate(200), vo.NewCoordinate(300), nil
	}

	sr.UpdateFunc = func(scooter *entity.Scooter) error {
		scooter.UpdateCoordinates(latitude, longitude)

		return errors.New("errors while updating scooter")
	}

	handler := event.NewTripStartedEventHandler(tr, sr, gl)
	err := handler.Handle(de)
	if err != nil {
		t.Fail()

		return
	}

	assert.Equals(t, latitude, scooter.Latitude)
	assert.Equals(t, longitude, scooter.Longitude)
}

func thenTripEventHandlerShouldNotReturnAnError(
	t *testing.T,
	tr *mock.TripRepository,
	sr *mock.ScooterRepository,
	gl *mock.GeographicLocator,
	de *domainEvent.DomainEvent,
) {
	latitude := vo.NewCoordinate(100)
	longitude := vo.NewCoordinate(200)

	glLatitude := vo.NewCoordinate(200)
	glLongitude := vo.NewCoordinate(300)

	scooter := new(entity.Scooter)
	scooter.UpdateCoordinates(latitude, longitude)

	trip := new(entity.Trip)
	trip.Scooter = scooter
	trip.Status = vo.NewFinishedStatus()

	tr.FindByUUIDFunc = func(u uuid.UUID) (*entity.Trip, error) {
		return trip, nil
	}

	gl.FindCurrentGeographicLocationFunc = func() (*vo.Coordinate, *vo.Coordinate, error) {
		return glLatitude, glLongitude, nil
	}

	sr.UpdateFunc = func(scooter *entity.Scooter) error {
		return nil
	}

	handler := event.NewTripStartedEventHandler(tr, sr, gl)
	err := handler.Handle(de)
	if err != nil {
		t.Fail()

		return
	}

	assert.Equals(t, glLatitude, scooter.Latitude)
	assert.Equals(t, glLongitude, scooter.Longitude)
}
