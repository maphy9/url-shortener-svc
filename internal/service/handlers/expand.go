package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

type Code string

func (c Code) isValid() bool {
	if len(c) == 0 {
		return false
	}
	for _, r := range c {
		if !(r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r >= '0' && r <='9') {
			return false
		}
	}
	return true
}

func ExpandURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	
	code := Code(chi.URLParam(r, "code"))
	if !code.isValid() {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	db := DB(r)
	query := Query{
		SQL: `
			SELECT url FROM url_mapping
			WHERE code = $1
		`,
		Args: []interface{} { code },
	}
	var url string
	if err := db.GetContext(ctx, &url, query); err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, url, http.StatusPermanentRedirect)
}