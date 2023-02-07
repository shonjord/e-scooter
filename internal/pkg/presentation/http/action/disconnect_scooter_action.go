package action

import (
	"encoding/json"

	"github.com/shonjord/e-scooter/internal/pkg/domain/command"
	"github.com/shonjord/e-scooter/internal/pkg/presentation/http"
)

type (
	scooterDisconnector interface {
		DisconnectScooter(scooter *command.DisconnectScooter) error
	}

	DisconnectScooterAction struct {
		disconnector scooterDisconnector
	}
)

// NewDisconnectScooterAction returns a new instance of this scooter disconnection action.
func NewDisconnectScooterAction(d scooterDisconnector) *DisconnectScooterAction {
	return &DisconnectScooterAction{
		disconnector: d,
	}
}

// Handle is responsible to disconnect a scooter and finish a trip.
func (a *DisconnectScooterAction) Handle(req *http.Request, res *http.Response) error {
	var cmd *command.DisconnectScooter

	if err := json.NewDecoder(req.Body()).Decode(&cmd); err != nil {
		return res.BadRequest(err)
	}

	return a.disconnector.DisconnectScooter(cmd)
}
