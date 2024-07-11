package operations

import (
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/truck/models"
)

type TruckDriverRemover struct {
	payload *models.TruckDriver
}

func (remover *TruckDriverRemover) TableName() string {
	return "truck_driver"
}

func (remover *TruckDriverRemover) Fields() []interface{} {
	return []interface{}{remover.payload.Id}
}

func (remover *TruckDriverRemover) Statement() string {
	return "delete from truck_driver where id = $1"
}

func (remover *TruckDriverRemover) Data() interface{} {
	return remover.payload
}

func NewTruckDriverRemoverCreator() *TruckDriverRemoverCreator {
	return &TruckDriverRemoverCreator{}
}

type TruckDriverRemoverCreator struct {
}

func (creator *TruckDriverRemoverCreator) Create(payload *models.TruckDriver) common.WriteOperation {
	return &TruckDriverRemover{
		payload: payload,
	}
}
