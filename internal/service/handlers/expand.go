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

	alias := chi.URLParam(r, "alias")
	matched, err := regexp.MatchString(`^[0-9a-zA-Z]{1,7}$`, alias)
	if err != nil {
		http.Error(w, fmt.Sprintf("Server error: %v", err), http.StatusInternalServerError)
		return
	}
	if !matched {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	aliasManager := AliasManager(r)
	url, err := aliasManager.GetOriginalUrl(ctx, alias)
	if err == sql.ErrNoRows {
		http.Error(w, "Not found", http.StatusNotFound)
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, url, http.StatusPermanentRedirect)
	}
}
