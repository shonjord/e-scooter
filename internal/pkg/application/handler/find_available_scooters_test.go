package handler_test

import (
	"errors"
	"testing"

	"github.com/shonjord/e-scooter/internal/pkg/application/handler"
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
	assert "github.com/shonjord/e-scooter/test/.assert"
	mock "github.com/shonjord/e-scooter/test/.mock"
)

// TestFindAvailableScooters responsible to test different important scenarios of this application handler.
func TestFindAvailableScooters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(
			*testing.T,
			*mock.ScooterRepository,
		)
	}{
		{
			// Given a call from the action to find available scooters
			scenario: "When scooter repository returns an errors",
			function: thenFindAvailableScootersShouldReturnScooterRepositoryError,
		},
		{
			// Given a call from the action to find available scooters
			scenario: "When handler return all available scooters successfully",
			function: thenFindAvailableScootersHandlerShouldReturnNoError,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(
				t,
				new(mock.ScooterRepository),
			)
		})
	}
}

func thenFindAvailableScootersShouldReturnScooterRepositoryError(
	t *testing.T,
	sr *mock.ScooterRepository,
) {
	sr.FindAvailableScootersFunc = func() ([]*entity.Scooter, error) {
		return nil, errors.New("mobile mobile")
	}

	h := handler.NewFindAvailableScooters(sr)
	_, err := h.FindAvailableScooters()

	assert.True(t, err != nil, "scooter repository should return an errors when something unexpected happens.")
}

func thenFindAvailableScootersHandlerShouldReturnNoError(
	t *testing.T,
	sr *mock.ScooterRepository,
) {
	sr.FindAvailableScootersFunc = func() ([]*entity.Scooter, error) {
		return []*entity.Scooter{new(entity.Scooter)}, nil
	}

	h := handler.NewFindAvailableScooters(sr)
	scooters, err := h.FindAvailableScooters()
	if err != nil {
		t.Fail()

		return
	}

	assert.True(t, len(scooters) > 0, "handler should return successfully all available scooters.")
}
