package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type ShortenURLRequest struct {
	URL string `json:"url"`
}

func (r ShortenURLRequest) isValid() bool {
	_, err := url.ParseRequestURI(r.URL)
	if err != nil {
		return false
	}
	return true
}

type Query struct {
	SQL string
	Args []interface{}
}

func (q Query) ToSql() (string, []interface{}, error) {
	return q.SQL, q.Args, nil
}

type URLMapping struct {
	URL string				`db:"url" json:"url"`
	Code string				`db:"code" json:"code"`
	CreatedAt time.Time		`db:"created_at" json:"created_at"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body ShortenURLRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Bad request body", http.StatusBadRequest)
		return
	}
	fmt.Println(body.URL)
	if !body.isValid() {
		http.Error(w, "Bad url", http.StatusBadRequest)
		return
	}

	db := DB(r)
	query := Query{
		SQL: `
			WITH ins AS (
				INSERT INTO url_mapping(url, code)
				VALUES($1, to_base62(nextval('code_sequence')))
				ON CONFLICT (url) DO NOTHING
				RETURNING code
			)
			SELECT code FROM ins
			UNION ALL
			SELECT code FROM url_mapping
			WHERE url = $1 LIMIT 1
		`,
		Args: []interface{} { body.URL },
	}
	var code string
	if err := db.GetContext(ctx, &code, query); err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "text/plain")
	shortURL := fmt.Sprintf("http://%s/%s", r.Host, code)
	w.Write([]byte(shortURL))
}