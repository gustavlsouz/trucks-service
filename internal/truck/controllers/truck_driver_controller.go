package controllers

import (
	"net/http"

	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/truck/models"
)

type truckDriverQueryCreator struct{}

func NewTruckDriverQueryCreator() *truckDriverQueryCreator {
	return &truckDriverQueryCreator{}
}

func (creator *truckDriverQueryCreator) Create(request *http.Request) *models.TruckDriver {
	return &models.TruckDriver{
		Id:      request.URL.Query().Get("id"),
		TruckId: request.URL.Query().Get("truckId"),
	}
}

type truckDriverDeleteCriteriaCreator struct{}

func NewTruckDriverDeleteCriteriaCreator() *truckDriverDeleteCriteriaCreator {
	return &truckDriverDeleteCriteriaCreator{}
}

func (creator *truckDriverDeleteCriteriaCreator) Create(request *http.Request) *models.TruckDriver {
	truckDriverId := request.URL.Query().Get("id")
	return &models.TruckDriver{Id: truckDriverId}
}

func NewTruckDriverController(
	reader common.ReadOperationCreator[models.TruckDriver],
	inserterCreator common.WriteOperationCreator[models.TruckDriverPayload],
	deleterCreator common.WriteOperationCreator[models.TruckDriver],
	updaterCreator common.WriteOperationCreator[models.TruckDriver],
) common.CrudController[models.TruckDriverPayload, models.TruckDriver, models.TruckDriver] {
	return common.NewCrudController(
		common.NewReader[models.TruckDriver, models.TruckDriver](reader),
		common.NewWriter(inserterCreator),
		common.NewWriter(updaterCreator),
		common.NewWriter(deleterCreator),
		NewTruckDriverQueryCreator(),
		NewTruckDriverDeleteCriteriaCreator(),
	)
}
