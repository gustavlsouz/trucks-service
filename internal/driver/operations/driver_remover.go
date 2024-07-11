package operations

import (
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/driver/models"
)

type DriverRemover struct {
	payload *models.Driver
}

func (remover *DriverRemover) TableName() string {
	return "driver"
}

func (remover *DriverRemover) Fields() []interface{} {
	return []interface{}{remover.payload.Id}
}

func (remover *DriverRemover) Statement() string {
	return "delete from driver where id = $1"
}

func (remover *DriverRemover) Data() interface{} {
	return remover.payload
}

func NewDriverRemoverCreator() *driverRemoverCreator {
	return &driverRemoverCreator{}
}

type driverRemoverCreator struct {
}

func (creator *driverRemoverCreator) Create(payload *models.Driver) common.WriteOperation {
	return &DriverRemover{
		payload: payload,
	}
}
