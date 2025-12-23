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
	res, err := json.Marshal(body)
	if err != nil {
		http.Error(w, "I am a teapot", http.StatusInternalServerError)
		return
	}
	w.Write(res)
}