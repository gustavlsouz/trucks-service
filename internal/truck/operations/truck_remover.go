package operations

import (
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/truck/models"
)

type TruckRemover struct {
	payload *models.Truck
}

func (remover *TruckRemover) TableName() string {
	return "truck"
}

func (remover *TruckRemover) Fields() []interface{} {
	return []interface{}{remover.payload.Id}
}

func (remover *TruckRemover) Statement() string {
	return "delete from truck where id = $1"
}

func (remover *TruckRemover) Data() interface{} {
	return remover.payload
}

func NewTruckRemoverCreator() *truckRemoverCreator {
	return &truckRemoverCreator{}
}

type truckRemoverCreator struct {
}

func (creator *truckRemoverCreator) Create(payload *models.Truck) common.WriteOperation {
	return &TruckRemover{
		payload: payload,
	}
}
