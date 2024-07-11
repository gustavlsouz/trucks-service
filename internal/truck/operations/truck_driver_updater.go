package operations

import (
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/truck/models"
)

type TruckDriverUpdater struct {
	payload *models.TruckDriver
}

func (updater *TruckDriverUpdater) TableName() string {
	return "truck_driver"
}

func (updater *TruckDriverUpdater) Fields() []interface{} {
	return []interface{}{updater.payload.Id, updater.payload.DriverId, updater.payload.CreatedAt}
}

func (updater *TruckDriverUpdater) Statement() string {
	return "update truck_driver set driverId = $2, createdAt = $3 where id = $1"
}

func (updater *TruckDriverUpdater) Data() interface{} {
	return updater.payload
}

func NewTruckDriverUpdaterCreator() *TruckDriverUpdaterCreator {
	return &TruckDriverUpdaterCreator{}
}

type TruckDriverUpdaterCreator struct{}

func (creator *TruckDriverUpdaterCreator) Create(payload *models.TruckDriver) common.WriteOperation {
	return &TruckDriverUpdater{
		payload: &models.TruckDriver{
			DriverId:  payload.DriverId,
			TruckId:   payload.TruckId,
			Id:        payload.Id,
			CreatedAt: payload.CreatedAt,
		},
	}
}
