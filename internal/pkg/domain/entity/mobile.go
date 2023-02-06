package entity

import (
	"time"

	"github.com/google/uuid"
)

type (
	Mobile struct {
		ID        int       `json:"-"`
		UUID      uuid.UUID `json:"uuid"`
		CreatedAt time.Time `json:"-"`
		UpdatedAt time.Time `json:"-"`
	}
)
