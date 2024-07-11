package common

import (
	"context"
	"fmt"
	"log"

	"github.com/gustavlsouz/trucks-service/internal/common/persistence"
)

type WriteOperation interface {
	TableName() string
	Fields() []interface{}
	Statement() string
	Data() interface{}
}

type WriteOperationCreator[T any] interface {
	Create(*T) WriteOperation
}

type RepositoryWriter[T any] interface {
	Write(context.Context, *T) (any, error)
}

type WriterService[T any] interface {
	Execute(ctx context.Context, model *T) (any, error)
}

func NewWriterService[T any](repository RepositoryWriter[T]) WriterService[T] {
	return &writerService[T]{
		repositoryWriter: repository,
	}
}

type writerService[T any] struct {
	repositoryWriter RepositoryWriter[T]
}

func (operator *writerService[T]) Execute(ctx context.Context, model *T) (any, error) {
	return operator.repositoryWriter.Write(ctx, model)
}

func NewRepositoryWriter[T any](operationCreator WriteOperationCreator[T]) RepositoryWriter[T] {
	return &repositoryWriter[T]{
		operationCreator: operationCreator,
		persistence:      persistence.GetPersistenceInstance(),
	}
}

type repositoryWriter[T any] struct {
	operationCreator WriteOperationCreator[T]
	persistence      persistence.Persistence
}

func (rWriter *repositoryWriter[T]) Write(ctx context.Context, model *T) (any, error) {
	operation := rWriter.operationCreator.Create(model)
	statment := operation.Statement()
	fields := operation.Fields()
	log.Println("changing:", operation.TableName(), "statment:", statment, "fields:", fields)
	_, err := rWriter.persistence.Database().ExecContext(ctx, statment, fields...)

	if err != nil {
		return nil, fmt.Errorf("error to execute statment: %w", err)
	}
	return operation.Data(), nil
}

func NewWriter[T any](operationCreator WriteOperationCreator[T]) WriterService[T] {
	return NewWriterService(
		NewRepositoryWriter(operationCreator),
	)
}
