// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package shelter

import (
	"context"
)

type Querier interface {
	AllDogs(ctx context.Context) ([]Dog, error)
	CreateDog(ctx context.Context, arg CreateDogParams) (Dog, error)
	DeleteDog(ctx context.Context, name string) error
	ReadDog(ctx context.Context, name string) (Dog, error)
	UpdateDog(ctx context.Context, arg UpdateDogParams) (Dog, error)
}

var _ Querier = (*Queries)(nil)
