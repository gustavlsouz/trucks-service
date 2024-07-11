package integration

import (
	"math/rand/v2"
	"strconv"
	"strings"
	"testing"

	"github.com/gustavlsouz/trucks-service/internal/truck/models"
)

func NewRandomPlateNumber() string {
	return strings.ToUpper(strconv.FormatInt(rand.Int64N(9999999), 16))
}

type TruckUpdater struct{}

func (truckUpdater *TruckUpdater) Update(truck models.Truck) models.Truck {
	truck.PlateNumber = NewRandomPlateNumber()
	return truck
}

type TruckComparer struct{}

func (truckComparer *TruckComparer) Compare(truckX models.Truck, truckY models.Truck) bool {
	return truckX.PlateNumber == truckY.PlateNumber
}

type TruckIdentifier struct{}

func (truckIdentifier *TruckIdentifier) Identify(truck models.Truck) string {
	return truck.Id
}

func NewMockedTruck() models.Truck {
	return models.Truck{PlateNumber: NewRandomPlateNumber()}
}

func TestTruck(t *testing.T) {

	httpUrl := "http://localhost:8080/api/truck"
	truckTest := NewMockedTruck()

	crudTesting := NewCrudTesting[models.Truck](
		"truck",
		&TruckUpdater{},
		&TruckComparer{},
		&TruckIdentifier{},
	)

	crudTesting.Execute(t, httpUrl, truckTest)

}
