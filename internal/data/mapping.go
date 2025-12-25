package data

import (
	"context"
	"time"
)

type MappingQ interface {
	GetByUrl(ctx context.Context, url string) (Mapping, error)

	GetByAlias(ctx context.Context, alias string) (Mapping, error)

	GetCode(ctx context.Context) (int64, error)

	Create(ctx context.Context, mapping Mapping) (Mapping, error)
}

type Mapping struct {
	Url       string    `db:"url" structs:"url"`
	Alias     string    `db:"alias" structs:"alias"`
	CreatedAt time.Time `db:"created_at" structs:"-"`
}
