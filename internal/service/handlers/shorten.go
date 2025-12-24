package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/maphy9/url-shortener-svc/internal/service/requests"
)

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var body requests.ShortenURLRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		http.Error(w, "Bad request body", http.StatusBadRequest)
		return
	}
	if !body.IsValid() {
		http.Error(w, "Bad url", http.StatusBadRequest)
		return
	}

	db := DB(r)
	query := SQLQuery{
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
		Args: []interface{}{body.URL},
	}
	var code string
	if err := db.GetContext(ctx, &code, query); err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	shortURL := fmt.Sprintf("%s://%s/%s", scheme, r.Host, code)

	log := Log(r)
	w.Header().Set("Content-Type", "text/plain")
	_, err := w.Write([]byte(shortURL))
	if err != nil {
		log.WithError(err).Error("Write failed")
		return
	}
}
