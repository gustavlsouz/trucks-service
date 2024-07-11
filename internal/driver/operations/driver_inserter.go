package operations

import (
	"time"

	"github.com/google/uuid"
	"github.com/gustavlsouz/trucks-service/internal/common"
	"github.com/gustavlsouz/trucks-service/internal/driver/models"
)

type DriverInserter struct {
	payload *models.Driver
}

func (inserter *DriverInserter) TableName() string {
	return "driver"
}

func (inserter *DriverInserter) Fields() []interface{} {
	return []interface{}{inserter.payload.Id, inserter.payload.Document, inserter.payload.CreatedAt}
}

func (inserter *DriverInserter) Statement() string {
	return "insert into driver (id, document, createdAt) values ($1, $2, $3)"
}

func (inserter *DriverInserter) Data() interface{} {
	return inserter.payload
}

func NewDriverInserterCreator() *DriverInserterCreator {
	return &DriverInserterCreator{}
}

type DriverInserterCreator struct{}

func (creator *DriverInserterCreator) Create(payload *models.DriverPayload) common.WriteOperation {
	return &DriverInserter{
		payload: &models.Driver{
			Id:        uuid.NewString(),
			Document:  payload.Document,
			CreatedAt: time.Now(),
		},
	}
}
