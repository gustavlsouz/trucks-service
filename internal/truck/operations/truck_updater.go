package operations

import (
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/truck/models"
)

type TruckUpdater struct {
	payload *models.Truck
}

func (inserter *TruckUpdater) TableName() string {
	return "truck"
}

func (inserter *TruckUpdater) Fields() []interface{} {
	return []interface{}{inserter.payload.Id, inserter.payload.PlateNumber, inserter.payload.CreatedAt}
}

func (inserter *TruckUpdater) Statement() string {
	return "update truck set plateNumber = $2, createdAt = $3 where id = $1"
}

func (inserter *TruckUpdater) Data() interface{} {
	return inserter.payload
}

func NewTruckUpdaterCreator() *TruckUpdaterCreator {
	return &TruckUpdaterCreator{}
}

type TruckUpdaterCreator struct{}

func (creator *TruckUpdaterCreator) Create(payload *models.Truck) common.WriteOperation {
	return &TruckUpdater{
		payload: &models.Truck{
			Id:          payload.Id,
			PlateNumber: payload.PlateNumber,
			CreatedAt:   payload.CreatedAt,
		},
	}
}
