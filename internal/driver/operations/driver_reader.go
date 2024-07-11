package operations

import (
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/driver/models"
)

type DriverReader struct {
	payload *models.Driver
}

func (reader *DriverReader) TableName() string {
	return "driver"
}

func (reader *DriverReader) Args() []interface{} {
	if reader.payload.Id != "" {
		return []interface{}{reader.payload.Id}
	}

	if reader.payload.Document != "" {
		return []interface{}{reader.payload.Document}
	}

	return []interface{}{}
}

func (reader *DriverReader) Query() string {
	if reader.payload.Id != "" {
		return "select id, document, createdAt from driver where id = $1"
	}

	if reader.payload.Document != "" {
		return "select id, document, createdAt from driver where document = $1"
	}

	// no pagination to simplify
	return "select id, document, createdAt from driver limit 100"
}

func NewDriverReaderCreator() *DriverReaderCreator {
	return &DriverReaderCreator{}
}

type DriverReaderCreator struct{}

func (creator *DriverReaderCreator) Create(payload *models.Driver) common.ReadOperation {
	return &DriverReader{
		payload: payload,
	}
}
