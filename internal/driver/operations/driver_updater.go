package operations

import (
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/driver/models"
)

type DriverUpdater struct {
	payload *models.Driver
}

func (updater *DriverUpdater) TableName() string {
	return "driver"
}

func (updater *DriverUpdater) Fields() []interface{} {
	return []interface{}{updater.payload.Id, updater.payload.Document, updater.payload.CreatedAt}
}

func (updater *DriverUpdater) Statement() string {
	return "update driver set document = $2, createdAt = $3 where id = $1"
}

func (updater *DriverUpdater) Data() interface{} {
	return updater.payload
}

func NewDriverUpdaterCreator() *DriverUpdaterCreator {
	return &DriverUpdaterCreator{}
}

type DriverUpdaterCreator struct{}

func (creator *DriverUpdaterCreator) Create(payload *models.Driver) common.WriteOperation {
	return &DriverUpdater{
		payload: &models.Driver{
			Id:        payload.Id,
			Document:  payload.Document,
			CreatedAt: payload.CreatedAt,
		},
	}
}
