package action

import (
	"github.com/shonjord/e-scooter/internal/pkg/domain/entity"
	"github.com/shonjord/e-scooter/internal/pkg/presentation/http"
)

type (
	scootersFinder interface {
		FindAvailableScooters() ([]*entity.Scooter, error)
	}

	GetAvailableScootersAction struct {
		finder scootersFinder
	}
)

// NewGetAvailableScootersAction returns a new instance of this action
// whose responsibility is to get available scooters.
func NewGetAvailableScootersAction(f scootersFinder) *GetAvailableScootersAction {
	return &GetAvailableScootersAction{
		finder: f,
	}
}

// Handle is responsible to get all available scooters in the system.
func (a *GetAvailableScootersAction) Handle(_ *http.Request, res *http.Response) error {
	scooters, err := a.finder.FindAvailableScooters()
	if err != nil {
		return res.InternalServerError(err)
	}

	if err = res.WriteStruct(scooters); err != nil {
		return res.InternalServerError(err)
	}

	return nil
}
