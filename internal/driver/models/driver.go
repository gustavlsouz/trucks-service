package models

import "time"

type Driver struct {
	Id        string    `json:"id,omitempty"`
	Document  string    `json:"document,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

type DriverPayload struct {
	Document string `json:"document"`
}
