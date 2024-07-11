package controllers

import (
	"net/http"

	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/driver/models"
)

type driverQueryCreator struct{}

func NewDriverQueryCreator() *driverQueryCreator {
	return &driverQueryCreator{}
}

func (creator *driverQueryCreator) Create(request *http.Request) *models.Driver {
	return &models.Driver{
		Id:       request.URL.Query().Get("id"),
		Document: request.URL.Query().Get("document"),
	}
}

type driverDeleteCriteriaCreator struct{}

func NewDriverDeleteCriteriaCreator() *driverDeleteCriteriaCreator {
	return &driverDeleteCriteriaCreator{}
}

func (creator *driverDeleteCriteriaCreator) Create(request *http.Request) *models.Driver {
	driverId := request.URL.Query().Get("id")
	return &models.Driver{Id: driverId}
}

func NewDriverController(
	reader common.ReadOperationCreator[models.Driver],
	inserterCreator common.WriteOperationCreator[models.DriverPayload],
	deleterCreator common.WriteOperationCreator[models.Driver],
	updaterCreator common.WriteOperationCreator[models.Driver],
) common.CrudController[models.DriverPayload, models.Driver, models.Driver] {
	return common.NewCrudController(
		common.NewReader[models.Driver, models.Driver](reader),
		common.NewWriter(inserterCreator),
		common.NewWriter(updaterCreator),
		common.NewWriter(deleterCreator),
		NewDriverQueryCreator(),
		NewDriverDeleteCriteriaCreator(),
	)
}
