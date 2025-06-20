// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0

package sqlc

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CreateCovid(ctx context.Context, arg CreateCovidParams) (Covid, error)
	DeleteCovid(ctx context.Context, id uuid.UUID) error
	GetCovid(ctx context.Context, id uuid.UUID) (Covid, error)
	GetCovids(ctx context.Context) ([]Covid, error)
	UpdateCovid(ctx context.Context, arg UpdateCovidParams) (Covid, error)
}

var _ Querier = (*Queries)(nil)
