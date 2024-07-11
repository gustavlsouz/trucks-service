package controllers

import (
	"net/http"

	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/truck/models"
)

type truckQueryCreator struct{}

func NewTruckQueryCreator() *truckQueryCreator {
	return &truckQueryCreator{}
}

func (creator *truckQueryCreator) Create(request *http.Request) *models.Truck {
	return &models.Truck{
		Id:          request.URL.Query().Get("id"),
		PlateNumber: request.URL.Query().Get("plateNumber"),
	}
}

type truckDeleteCriteriaCreator struct{}

func NewTruckDeleteCriteriaCreator() *truckDeleteCriteriaCreator {
	return &truckDeleteCriteriaCreator{}
}

func (creator *truckDeleteCriteriaCreator) Create(request *http.Request) *models.Truck {
	truckId := request.URL.Query().Get("id")
	return &models.Truck{Id: truckId}
}

func NewTruckController(
	reader common.ReadOperationCreator[models.Truck],
	inserterCreator common.WriteOperationCreator[models.TruckPayload],
	deleterCreator common.WriteOperationCreator[models.Truck],
	updaterCreator common.WriteOperationCreator[models.Truck],
) common.CrudController[models.TruckPayload, models.Truck, models.Truck] {
	return common.NewCrudController(
		common.NewReader[models.Truck, models.Truck](reader),
		common.NewWriter(inserterCreator),
		common.NewWriter(updaterCreator),
		common.NewWriter(deleterCreator),
		NewTruckQueryCreator(),
		NewTruckDeleteCriteriaCreator(),
	)
}
