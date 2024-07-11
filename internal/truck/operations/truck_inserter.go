package operations

import (
	"time"

	"github.com/google/uuid"
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/truck/models"
)

type TruckInserter struct {
	payload *models.Truck
}

func (inserter *TruckInserter) TableName() string {
	return "truck"
}

func (inserter *TruckInserter) Fields() []interface{} {
	return []interface{}{inserter.payload.Id, inserter.payload.PlateNumber, inserter.payload.CreatedAt}
}

func (inserter *TruckInserter) Statement() string {
	return "insert into truck (id, plateNumber, createdAt) values ($1, $2, $3)"
}

func (inserter *TruckInserter) Data() interface{} {
	return inserter.payload
}

func NewTruckInserterCreator() *TruckInserterCreator {
	return &TruckInserterCreator{}
}

type TruckInserterCreator struct{}

func (creator *TruckInserterCreator) Create(payload *models.TruckPayload) common.WriteOperation {
	return &TruckInserter{
		payload: &models.Truck{
			Id:          uuid.NewString(),
			PlateNumber: payload.PlateNumber,
			CreatedAt:   time.Now(),
		},
	}
}
