package operations

import (
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/truck/models"
)

type TruckDriverReader struct {
	payload *models.TruckDriver
}

func (reader *TruckDriverReader) TableName() string {
	return "truck_driver"
}

func (reader *TruckDriverReader) Args() []interface{} {
	if reader.payload.Id != "" {
		return []interface{}{reader.payload.Id}
	}

	if reader.payload.TruckId != "" {
		return []interface{}{reader.payload.TruckId}
	}

	return []interface{}{}
}

func (reader *TruckDriverReader) Query() string {
	if reader.payload.Id != "" {
		return "select driverId, truckId, id, createdAt from truck_driver where id = $1"
	}

	if reader.payload.TruckId != "" {
		return "select driverId, truckId, id, createdAt from truck_driver where truckId = $1"
	}

	// no pagination to simplify
	return "select driverId, truckId, id, createdAt from truck_driver limit 100"
}

func NewTruckDriverReaderCreator() *TruckDriverReaderCreator {
	return &TruckDriverReaderCreator{}
}

type TruckDriverReaderCreator struct{}

func (creator *TruckDriverReaderCreator) Create(payload *models.TruckDriver) common.ReadOperation {
	return &TruckDriverReader{
		payload: payload,
	}
}
