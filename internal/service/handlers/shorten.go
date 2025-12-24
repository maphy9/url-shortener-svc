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

	aliasManager := AliasManager(r)
	alias, err := aliasManager.GetAlias(ctx, body.Url)
	if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	shortUrl := fmt.Sprintf("%s://%s/%s", scheme, r.Host, alias)

	log := Log(r)
	w.Header().Set("Content-Type", "text/plain")
	if _, err := w.Write([]byte(shortUrl)); err != nil {
		log.WithError(err).Error("Write failed")
		return
	}
}
