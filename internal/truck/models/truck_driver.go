package models

import "time"

type TruckDriverPayload struct {
	DriverId string `json:"driverId"`
	TruckId  string `json:"truckId"`
}

type TruckDriver struct {
	DriverId  string    `json:"driverId,omitempty"`
	TruckId   string    `json:"truckId,omitempty"`
	Id        string    `json:"id,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
