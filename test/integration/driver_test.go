package integration

import (
	"math/rand/v2"
	"strconv"
	"strings"
	"testing"

	"github.com/gustavlsouz/trucks-service/internal/driver/models"
)

type DriverUpdater struct{}

func randomDocument() string {
	return strings.ToUpper(strconv.FormatInt(rand.Int64N(99999999999), 16))
}

func (driverUpdater *DriverUpdater) Update(driver models.Driver) models.Driver {
	driver.Document = randomDocument()
	return driver
}

type DriverComparer struct{}

func (driverComparer *DriverComparer) Compare(driverX models.Driver, driverY models.Driver) bool {
	return driverX.Document == driverY.Document
}

type DriverIdentifier struct{}

func (driverIdentifier *DriverIdentifier) Identify(driver models.Driver) string {
	return driver.Id
}

func NewMockedDriver() models.Driver {
	return models.Driver{Document: randomDocument()}
}

func TestDriver(t *testing.T) {

	httpUrl := "http://localhost:8080/api/driver"
	driverTest := NewMockedDriver()

	crudTesting := NewCrudTesting(
		"driver",
		&DriverUpdater{},
		&DriverComparer{},
		&DriverIdentifier{},
	)

	crudTesting.Execute(t, httpUrl, driverTest)

}
