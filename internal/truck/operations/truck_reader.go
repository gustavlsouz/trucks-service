package operations

import (
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/truck/models"
)

type TruckReader struct {
	payload *models.Truck
}

func (reader *TruckReader) TableName() string {
	return "truck"
}

func (reader *TruckReader) Args() []interface{} {
	if reader.payload.Id != "" {
		return []interface{}{reader.payload.Id}
	}

	if reader.payload.PlateNumber != "" {
		return []interface{}{reader.payload.PlateNumber}
	}

	return []interface{}{}
}

func (reader *TruckReader) Query() string {
	if reader.payload.Id != "" {
		return "select id, plateNumber, createdAt from truck where id = $1"
	}

	if reader.payload.PlateNumber != "" {
		return "select id, plateNumber, createdAt from truck where plateNumber = $1"
	}

	// no pagination to simplify
	return "select id, plateNumber, createdAt from truck limit 100"
}

func NewTruckReaderCreator() *TruckReaderCreator {
	return &TruckReaderCreator{}
}

type TruckReaderCreator struct{}

func (creator *TruckReaderCreator) Create(payload *models.Truck) common.ReadOperation {
	return &TruckReader{
		payload: payload,
	}
}
