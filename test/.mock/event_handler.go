package mock

import "github.com/shonjord/e-scooter/internal/pkg/domain/event"

type (
	EventHandler struct {
		HandleFunc       func(*event.DomainEvent) error
		SubscribedToFunc func() string
	}
)

// Handle refer to the consumer of the interface for documentation.
func (m *EventHandler) Handle(e *event.DomainEvent) error {
	return m.HandleFunc(e)
}

// SubscribedTo refer to the consumer of the interface for documentation.
func (m *EventHandler) SubscribedTo() string {
	return m.SubscribedToFunc()
}
