package models

import "time"

type TruckPayload struct {
	PlateNumber string `json:"plateNumber"`
}

type Truck struct {
	Id          string    `json:"id,omitempty"`
	PlateNumber string    `json:"plateNumber,omitempty"`
	CreatedAt   time.Time `json:"createdAt,omitempty"`
}
