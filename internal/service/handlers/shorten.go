package handlers

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"net/url"
	"time"
)

type ShortenURLRequest struct {
	URL string `json:"url"`
}

func (r ShortenURLRequest) IsValid() bool {
    u, err := url.Parse(r.URL)
	if u.Scheme == "" {
		u.Scheme = "https"
	}
    return err == nil && u.Host != ""
}

type Query struct {
	SQL string
	Args []interface{}
}

func (q Query) ToSql() (string, []interface{}, error) {
	return q.SQL, q.Args, nil
}

type URLMapping struct {
	OriginalURL string		`db:"original_url" json:"original_url"`
	ShortenedURL string		`db:"shortened_url" json:"shortened_url"`
	CreatedAt time.Time		`db:"created_at" json:"created_at"`
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	var body ShortenURLRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Bad request body", http.StatusBadRequest)
		return
	}	
	if !body.IsValid() {
		http.Error(w, "Bad url", http.StatusBadRequest)
		return
	}
	db := DB(r)
	var shortenedUrl string
	if err := db.Get(&shortenedUrl, Query{
		SQL: `SELECT shortened_url FROM url_mapping
		WHERE original_url = $1;`,
		Args: []interface{} { body.URL },
	}); err == nil {
		w.Write([]byte("http://localhost" + r.RequestURI + "/" + shortenedUrl))
		return
	}

	// TODO: Refactor numbering logic
	var countString string
	if err := db.Get(
		&countString,
		Query{SQL: "SELECT COUNT(*) FROM url_mapping"},
	); err != nil {
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
		return
	}
	count := new(big.Int)
	count, _ = count.SetString(countString, 10)
	if count == nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	count.Add(count, new(big.Int).SetInt64(1))
	base62 := count.Text(62)
	if err := db.Exec(Query{
		SQL: `INSERT INTO url_mapping(
			original_url, shortened_url
		) VALUES($1, $2);`,
		Args: []interface{} {body.URL, base62},
	}); err != nil {
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
		return
	}
	w.Write([]byte("http://localhost" + r.RequestURI + "/" + base62))
}