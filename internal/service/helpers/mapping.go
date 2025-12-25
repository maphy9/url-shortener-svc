package helpers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/maphy9/url-shortener-svc/internal/data"
	"github.com/maphy9/url-shortener-svc/internal/service/util"
)

func GetShortUrl(r *http.Request, url string) (string, error) {
	ctx := r.Context()
	masterQ := DB(r)

	var alias string
	transaction := func(q data.MasterQ) error {
		mapping, err := q.Mapping().GetByUrl(ctx, url)
		if err == nil {
			alias = mapping.Alias
			return nil
		}
		if err != sql.ErrNoRows {
			return err
		}

		code, err := q.Mapping().GetCode(ctx)
		if err != nil {
			return err
		}

		mapping, err = q.Mapping().Create(ctx, data.Mapping{
			Url: url,
			Alias: util.ToBase62(code),
		})
		if err != nil {
			return err
		}
		alias = mapping.Alias
		return nil
	}
	err := masterQ.Transaction(transaction)

	scheme := "http"
	if r.TLS != nil {
		scheme = "https"
	}
	shortUrl := fmt.Sprintf("%s://%s/%s", scheme, r.Host, alias)
	return shortUrl, err
}

func GetOriginalUrl(r *http.Request, alias string) (string, error) {
	ctx := r.Context()
	masterQ := DB(r)

	mapping, err := masterQ.Mapping().GetByAlias(ctx, alias)
	if err != nil {
		return "", err
	}
	return mapping.Url, nil
}