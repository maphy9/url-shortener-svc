package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func ExpandURL(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// TODO: Uncomment validation when aliasing algorithm was changed
	alias := chi.URLParam(r, "alias")
	// matched, err := regexp.MatchString(`^[0-9a-zA-Z]{1,7}$`, alias)
	// if err != nil {
	// 	http.Error(w, fmt.Sprintf("Server error: %v", err), http.StatusInternalServerError)
	// 	return
	// }
	// if !matched {
	// 	http.Error(w, "Not found", http.StatusNotFound)
	// 	return
	// }

	mappingQ := DB(r)
	mapping, err := mappingQ.Mapping().GetByAlias(ctx, alias)
	if err == sql.ErrNoRows {
		http.Error(w, "Not found", http.StatusNotFound)
	} else if err != nil {
		http.Error(w, fmt.Sprintf("Database error: %v", err), http.StatusInternalServerError)
	} else {
		http.Redirect(w, r, mapping.Url, http.StatusPermanentRedirect)
	}
}
