package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
)

func ExpandURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	code := chi.URLParam(r, "code")
	matched, err := regexp.MatchString(`^[0-9a-zA-Z]{1,7}$`, code)
	if err != nil {
		http.Error(w, fmt.Sprintf("Server error: %v", err), http.StatusInternalServerError)
		return		
	}
	if !matched {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	db := DB(r)
	query := SQLQuery{
		SQL: `
			SELECT url FROM url_mapping
			WHERE code = $1
		`,
		Args: []interface{}{code},
	}
	var url string
	if err := db.GetContext(ctx, &url, query); err == sql.ErrNoRows {
		http.Error(w, "Not found", http.StatusNotFound)
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	}
}
