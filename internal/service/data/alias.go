package data

import (
	"context"

	"github.com/Masterminds/squirrel"
	"gitlab.com/distributed_lab/kit/pgdb"
)

type AliasManager interface {
	GetAlias(ctx context.Context, url string) (string, error)
	GetOriginalUrl(ctx context.Context, alias string) (string, error)
}

type aliasManager struct {
	db  *pgdb.DB
	sql squirrel.StatementBuilderType
}

func NewUrlAliasesManager(db *pgdb.DB) AliasManager {
	return &aliasManager{
		db:  db,
		sql: squirrel.StatementBuilder,
	}
}

func (m *aliasManager) GetAlias(ctx context.Context, url string) (string, error) {
	var result string
	query := m.sql.Select().
		Column("get_alias(?)", url).
		PlaceholderFormat(squirrel.Dollar)
	err := m.db.GetContext(ctx, &result, query)
	return result, err
}

func (m *aliasManager) GetOriginalUrl(ctx context.Context, alias string) (string, error) {
	var result string
	query := m.sql.Select().
		Column("get_url(?)", alias).
		PlaceholderFormat(squirrel.Dollar)
	err := m.db.GetContext(ctx, &result, query)
	return result, err
}
