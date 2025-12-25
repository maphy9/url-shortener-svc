package pg

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/fatih/structs"
	"github.com/maphy9/url-shortener-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

const mappingTableName = "mapping"

func newMappingQ(db *pgdb.DB) data.MappingQ {
	return &mappingQ{
		db:  db,
		sql: squirrel.StatementBuilder,
	}
}

type mappingQ struct {
	db  *pgdb.DB
	sql squirrel.StatementBuilderType
}

func (m *mappingQ) GetByUrl(ctx context.Context, url string) (data.Mapping, error) {
	query := m.sql.Select("*").
		From(mappingTableName).
		Where("url = ?", url).
		PlaceholderFormat(squirrel.Dollar)

	var result data.Mapping
	err := m.db.GetContext(ctx, &result, query)
	return result, err
}

func (m *mappingQ) GetByAlias(ctx context.Context, alias string) (data.Mapping, error) {
	query := m.sql.Select("*").
		From(mappingTableName).
		Where("alias = ?", alias).
		PlaceholderFormat(squirrel.Dollar)

	var result data.Mapping
	err := m.db.GetContext(ctx, &result, query)
	return result, err
}

func (m *mappingQ) Create(ctx context.Context, mapping data.Mapping) (data.Mapping, error) {
	clauses := structs.Map(mapping)

	query := m.sql.Insert(mappingTableName).
		SetMap(clauses).
		Suffix("RETURNING *")

	var result data.Mapping
	err := m.db.GetContext(ctx, &result, query)
	return result, err
}
