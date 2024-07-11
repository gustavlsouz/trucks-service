package operations

import (
	"time"

	"github.com/google/uuid"
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/truck/models"
)

type TruckDriverInserter struct {
	payload *models.TruckDriver
}

func (inserter *TruckDriverInserter) TableName() string {
	return "truck_driver"
}

func (inserter *TruckDriverInserter) Fields() []interface{} {
	return []interface{}{inserter.payload.Id, inserter.payload.TruckId, inserter.payload.DriverId, inserter.payload.CreatedAt}
}

func (inserter *TruckDriverInserter) Statement() string {
	return "insert into truck_driver (id, truckId, driverId, createdAt) values ($1, $2, $3, $4)"
}

func (inserter *TruckDriverInserter) Data() interface{} {
	return inserter.payload
}

func NewTruckDriverInserterCreator() *TruckDriverInserterCreator {
	return &TruckDriverInserterCreator{}
}

type TruckDriverInserterCreator struct{}

func (creator *TruckDriverInserterCreator) Create(payload *models.TruckDriverPayload) common.WriteOperation {
	return &TruckDriverInserter{
		payload: &models.TruckDriver{
			Id:        uuid.NewString(),
			CreatedAt: time.Now(),
			DriverId:  payload.DriverId,
			TruckId:   payload.TruckId,
		},
	}
}
