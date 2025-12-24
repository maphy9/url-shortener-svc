package handlers

import "time"

type SQLQuery struct {
	SQL  string
	Args []interface{}
}

func (q SQLQuery) ToSql() (string, []interface{}, error) {
	return q.SQL, q.Args, nil
}

type URLMapping struct {
	URL       string    `db:"url" json:"url"`
	Code      string    `db:"code" json:"code"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
}
