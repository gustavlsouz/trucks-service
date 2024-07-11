package common

import "context"

type ReadUseCase interface {
	Execute(ctx context.Context) error
}
