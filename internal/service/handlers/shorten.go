package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
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

func (r ShortenURLRequest) ToSql() (string, []interface{}, error) {
	return fmt.Sprintf(`
		SELECT COUNT(*)
		FROM URL_MAPPING
		WHERE ORIGINAL_URL='%s';`, r.URL), []interface{}{}, nil
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL: %v, Method: %v\n", r.URL, r.Method)
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
	res, err := db.ExecWithResult(body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		http.Error(w, fmt.Sprintf("Internal server error: %v", err), http.StatusInternalServerError)
	}
	w.Write([]byte(fmt.Sprintf("Rows affected: %d", rowsAffected)))
}