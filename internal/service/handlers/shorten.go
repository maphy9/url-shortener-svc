package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand/v2"
	"net/http"

	"github.com/maphy9/url-shortener-svc/internal/data"
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

	// TODO: This is not concurrency safe, need to handle simulanious INSERT
	var alias string
	masterQ := DB(r)
	transaction := func(q data.MasterQ) error {
		mapping, err := q.Mapping().GetByUrl(ctx, body.Url)
		if err == nil {
			alias = mapping.Alias
			return nil
		}
		if err != sql.ErrNoRows {
			return err
		}
		// TODO: Change aliasing algorithm
		mapping, err = q.Mapping().Create(ctx, data.Mapping{
			Url: body.Url,
			Alias: fmt.Sprintf("%v", rand.Int64()),
		})
		if err != nil {
			return err
		}
		alias = mapping.Alias
		return nil
	}
	masterQ.Transaction(transaction)

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
