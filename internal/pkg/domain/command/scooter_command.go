package command

import "github.com/google/uuid"

type (
	ConnectScooter struct {
		MobileUUID  uuid.UUID `json:"mobile_uuid"`
		ScooterUUID uuid.UUID `json:"scooter_uuid"`
	}

	DisconnectScooter struct {
		MobileUUID  uuid.UUID `json:"mobile_uuid"`
		ScooterUUID uuid.UUID `json:"scooter_uuid"`
	}
)
