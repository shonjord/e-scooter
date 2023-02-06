package action

import (
	"encoding/json"

	"github.com/shonjord/e-scooter/internal/pkg/domain/command"
	"github.com/shonjord/e-scooter/internal/pkg/presentation/http"
)

type (
	scooterConnector interface {
		ConnectScooter(*command.ConnectScooter) error
	}

	ConnectScooterAction struct {
		connector scooterConnector
	}
)

// NewConnectScooterAction returns a new instance of a mobile connection to a scooter.
func NewConnectScooterAction(c scooterConnector) *ConnectScooterAction {
	return &ConnectScooterAction{
		connector: c,
	}
}

// Handle is responsible to make a connection between a mobile and a scooter for a new ride :).
func (a *ConnectScooterAction) Handle(req *http.Request, res *http.Response) error {
	var cmd *command.ConnectScooter

	if err := json.NewDecoder(req.Body()).Decode(&cmd); err != nil {
		return res.BadRequest(err)
	}

	return a.connector.ConnectScooter(cmd)
}
