package integration

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"testing"

	driverModels "github.com/gustavlsouz/trucks-service/internal/driver/models"
	"github.com/gustavlsouz/trucks-service/internal/truck/models"
	"gotest.tools/v3/assert"
)

func TestTruckDriver(t *testing.T) {

	httpUrl := "http://localhost:8080/api/truck/relation"
	client := &http.Client{}

	truck := createTruck(t)
	driver := createDriver(t)
	truckDriverTest := models.TruckDriver{
		DriverId: driver.Id,
		TruckId:  truck.Id,
	}

	t.Run("should create truck driver relation successfully", func(t *testing.T) {

		method := "POST"

		content, err := json.Marshal(truckDriverTest)
		assert.NilError(t, err)

		payload := bytes.NewReader(content)

		req, err := http.NewRequest(method, httpUrl, payload)

		assert.NilError(t, err)

		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		assert.NilError(t, err)
		defer res.Body.Close()

		err = ReaderToStruct(res.Body, &truckDriverTest)
		assert.NilError(t, err)

		assert.Equal(t, http.StatusCreated, res.StatusCode)
	})

	t.Run("should read truck driver relation successfully", func(t *testing.T) {

		method := "GET"

		req, err := http.NewRequest(method, httpUrl, nil)

		assert.NilError(t, err)

		res, err := client.Do(req)
		assert.NilError(t, err)
		defer res.Body.Close()

		var truckDrivers []models.TruckDriver
		err = ReaderToStruct(res.Body, &truckDrivers)
		assert.NilError(t, err)
		assert.Equal(t, true, len(truckDrivers) <= 100)

		var foundTruckDriver bool
		for _, truckDriver := range truckDrivers {
			if truckDriver.Id == truckDriverTest.Id {
				foundTruckDriver = true
			}
		}

		assert.Equal(t, true, foundTruckDriver)

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("should update truck driver relation successfully", func(t *testing.T) {

		newDriver := createDriver(t)

		method := "PUT"

		truckDriverUpdate := truckDriverTest

		truckDriverUpdate.DriverId = newDriver.Id

		content, err := json.Marshal(truckDriverUpdate)
		assert.NilError(t, err)

		payload := bytes.NewReader(content)

		req, err := http.NewRequest(method, httpUrl, payload)

		assert.NilError(t, err)

		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		assert.NilError(t, err)
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})

	t.Run("should delete truck driver relation successfully", func(t *testing.T) {

		method := "DELETE"

		parsedUrl, err := url.Parse(httpUrl)
		assert.NilError(t, err)

		query := parsedUrl.Query()
		query.Add("id", truckDriverTest.Id)
		parsedUrl.RawQuery = query.Encode()

		req, err := http.NewRequest(method, parsedUrl.String(), nil)

		assert.NilError(t, err)

		req.Header.Add("Content-Type", "application/json")

		res, err := client.Do(req)
		assert.NilError(t, err)
		defer res.Body.Close()

		assert.Equal(t, http.StatusOK, res.StatusCode)
	})
}

func createTruck(t *testing.T) models.Truck {
	method := "POST"
	httpUrl := "http://localhost:8080/api/truck"
	client := &http.Client{}

	content, err := json.Marshal(NewMockedTruck())
	assert.NilError(t, err)

	payload := bytes.NewReader(content)

	req, err := http.NewRequest(method, httpUrl, payload)

	assert.NilError(t, err)

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	assert.NilError(t, err)
	defer res.Body.Close()

	var truck models.Truck

	err = ReaderToStruct(res.Body, &truck)
	assert.NilError(t, err)

	return truck
}

func createDriver(t *testing.T) driverModels.Driver {
	method := "POST"
	httpUrl := "http://localhost:8080/api/driver"
	client := &http.Client{}

	content, err := json.Marshal(NewMockedDriver())
	assert.NilError(t, err)

	payload := bytes.NewReader(content)

	req, err := http.NewRequest(method, httpUrl, payload)

	assert.NilError(t, err)

	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	assert.NilError(t, err)
	defer res.Body.Close()

	var driver driverModels.Driver

	err = ReaderToStruct(res.Body, &driver)
	assert.NilError(t, err)

	return driver
}
