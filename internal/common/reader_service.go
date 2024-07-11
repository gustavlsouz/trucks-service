package common

import (
	"context"
	"fmt"
	"log"

	"github.com/gustavlsouz/trucks-service/internal/common/persistence"
)

type ReadOperation interface {
	TableName() string
	Args() []interface{}
	Query() string
}

type ReadOperationCreator[Q any] interface {
	Create(*Q) ReadOperation
}

type RepositoryReader[Q any, T any] interface {
	Read(context.Context, *Q) ([]T, error)
}

type ReaderService[Q any, T any] interface {
	Execute(ctx context.Context, model *Q) ([]T, error)
}

func NewReaderService[Q any, T any](repository RepositoryReader[Q, T]) ReaderService[Q, T] {
	return &readerService[Q, T]{
		repositoryWriter: repository,
	}
}

type readerService[Q any, T any] struct {
	repositoryWriter RepositoryReader[Q, T]
}

func (operator *readerService[Q, T]) Execute(ctx context.Context, model *Q) ([]T, error) {
	return operator.repositoryWriter.Read(ctx, model)
}

func NewRepositoryReader[Q any, T any](operationCreator ReadOperationCreator[Q]) RepositoryReader[Q, T] {
	return &repositoryReader[Q, T]{
		operationCreator: operationCreator,
		persistence:      persistence.GetPersistenceInstance(),
	}
}

type repositoryReader[Q any, T any] struct {
	operationCreator ReadOperationCreator[Q]
	persistence      persistence.Persistence
}

func (rReader *repositoryReader[Q, T]) Read(ctx context.Context, model *Q) ([]T, error) {
	operation := rReader.operationCreator.Create(model)
	query := operation.Query()
	args := operation.Args()
	log.Println("reading:", operation.TableName(), "query:", query, "args:", args)
	rows, err := rReader.persistence.Database().QueryContext(ctx, query, args...)

	if err != nil {
		return nil, fmt.Errorf("error to query: %w", err)
	}

	list := make([]T, 0)

	for rows.Next() {
		var item T
		err = persistence.GetPersistenceInstance().ScanStruct(rows, &item)

		if err != nil {
			return nil, fmt.Errorf("error to scan struct: %w", err)
		}

		list = append(list, item)
	}

	if err := rows.Close(); err != nil {
		return nil, fmt.Errorf("error to close rows: %w", err)
	}

	return list, nil
}

func NewReader[Q any, T any](operationCreator ReadOperationCreator[Q]) ReaderService[Q, T] {
	return NewReaderService(
		NewRepositoryReader[Q, T](operationCreator),
	)
}
