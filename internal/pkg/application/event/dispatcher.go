package event

import (
	"fmt"

	"github.com/shonjord/e-scooter/internal/pkg/domain/event"
	log "github.com/sirupsen/logrus"
)

type (
	eventHandler interface {
		Handle(*event.DomainEvent) error
		SubscribedTo() string
	}

	DomainEventDispatcher struct {
		handlers map[string]eventHandler
	}
)

// NewDomainEventDispatcher returns an event dispatcher responsible to allocate the handler and handle the event.
func NewDomainEventDispatcher(h ...interface{}) (*DomainEventDispatcher, error) {
	handlers := make(map[string]eventHandler, 0)

	for _, eh := range h {
		if handler, ok := eh.(eventHandler); ok {
			handlers[handler.SubscribedTo()] = handler

			continue
		}

		return nil, fmt.Errorf("event handler: %T, does not implement event handler interface", eh)
	}

	return &DomainEventDispatcher{handlers: handlers}, nil
}

// Dispatch responsible to dispatch the event to the registered handler.
func (e *DomainEventDispatcher) Dispatch(event *event.DomainEvent) error {
	handler, ok := e.handlers[event.Name]
	if !ok {
		return fmt.Errorf("handler for event: %s, doesn't exist", event.Name)
	}

	if !event.Async {
		return handler.Handle(event)
	}

	go func() {
		if err := handler.Handle(event); err != nil {
			log.WithError(err).Error("error while handling event")
		}
	}()

	return nil
}
