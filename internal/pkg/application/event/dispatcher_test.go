package event_test

import (
	"errors"
	"testing"

	"github.com/google/uuid"
	"github.com/shonjord/e-scooter/internal/pkg/application/event"
	domainEvent "github.com/shonjord/e-scooter/internal/pkg/domain/event"
	assert "github.com/shonjord/e-scooter/test/.assert"
	mock "github.com/shonjord/e-scooter/test/.mock"
)

// TestEventDispatcher responsible to test different important scenarios of application's event dispatcher.
func TestEventDispatcher(t *testing.T) {
	t.Parallel()

	tests := []struct {
		scenario string
		function func(*testing.T, *mock.EventHandler)
	}{
		{
			// Given a dependency that doesn't implement event handler interface
			scenario: "When instantiation of dispatcher with event handlers no implementing event handler interface",
			function: thenDispatcherShouldReturnError,
		},
		{
			// Given a valid domain event
			scenario: "When no handler is available to process the event",
			function: thenDispatcherShouldReturnNoHandlerFoundError,
		},
		{
			// Given a valid domain event
			scenario: "When event handler returns an errors",
			function: thenDispatcherShouldReturnEventHandlerError,
		},
		{
			// Given a valid domain event
			scenario: "When event handler doesn't return an errors",
			function: thenDispatcherShouldReturnNoError,
		},
	}

	for _, test := range tests {
		t.Run(test.scenario, func(t *testing.T) {
			test.function(t, new(mock.EventHandler))
		})
	}
}

func thenDispatcherShouldReturnError(t *testing.T, _ *mock.EventHandler) {
	_, err := event.NewDomainEventDispatcher([]interface{}{"asd"})

	assert.True(t, err != nil, "dispatcher should return errors when event handler doesn't implement interface.")
}

func thenDispatcherShouldReturnNoHandlerFoundError(t *testing.T, eh *mock.EventHandler) {
	eh.SubscribedToFunc = func() string {
		return "event"
	}

	dispatcher, err := event.NewDomainEventDispatcher(eh)
	if err != nil {
		t.Fail()

		return
	}

	err = dispatcher.Dispatch(&domainEvent.DomainEvent{
		UUID:  uuid.New(),
		Name:  "unknownEvent",
		Async: false,
	})

	assert.True(t, err != nil, "dispatcher should return errors when handler can't be found.")
}

func thenDispatcherShouldReturnEventHandlerError(t *testing.T, eh *mock.EventHandler) {
	eh.SubscribedToFunc = func() string {
		return "event"
	}

	eh.HandleFunc = func(e *domainEvent.DomainEvent) error {
		return errors.New("event handler errors")
	}

	dispatcher, err := event.NewDomainEventDispatcher(eh)
	if err != nil {
		t.Fail()

		return
	}

	err = dispatcher.Dispatch(&domainEvent.DomainEvent{
		UUID:  uuid.New(),
		Name:  "event",
		Async: false,
	})

	assert.True(t, err != nil, "dispatcher should return errors when handler can't be found.")
}

func thenDispatcherShouldReturnNoError(t *testing.T, eh *mock.EventHandler) {
	eh.SubscribedToFunc = func() string {
		return "event"
	}

	eh.HandleFunc = func(e *domainEvent.DomainEvent) error {
		return nil
	}

	dispatcher, err := event.NewDomainEventDispatcher(eh)
	if err != nil {
		t.Fail()

		return
	}

	err = dispatcher.Dispatch(&domainEvent.DomainEvent{
		UUID:  uuid.New(),
		Name:  "event",
		Async: false,
	})

	assert.True(t, err == nil, "dispatcher should not return errors when handler runs successfully.")
}
