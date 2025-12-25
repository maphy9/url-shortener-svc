package pg

import (
	"github.com/maphy9/url-shortener-svc/internal/data"
	"gitlab.com/distributed_lab/kit/pgdb"
)

func NewMasterQ(db *pgdb.DB) data.MasterQ {
	return &masterQ{db}
}

type masterQ struct {
	db *pgdb.DB
}

func (m *masterQ) Mapping() data.MappingQ {
	return newMappingQ(m.db)
}

func (m *masterQ) Transaction(fn func(q data.MasterQ) error) error {
	return m.db.Transaction(func() error {
		return fn(m)
	})
}