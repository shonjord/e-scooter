package mock

import "github.com/shonjord/e-scooter/internal/pkg/domain/event"

type (
	EventDispatcher struct {
		DispatchFunc func(*event.DomainEvent) error
	}
)

// Dispatch refer to the consumer of the interface for documentation.
func (m *EventDispatcher) Dispatch(e *event.DomainEvent) error {
	return m.DispatchFunc(e)
}
