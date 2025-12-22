package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ShortenURLRequest struct {
	URL string `json:"url"`
}

func (r ShortenURLRequest) String() string {
	return fmt.Sprintf("URL: %s", r.URL)
}

func ShortenURL(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("URL: %v, Method: %v\n", r.URL, r.Method)
	var body ShortenURLRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		fmt.Errorf("%v\n", err)
		return
	}
	w.Write([]byte(body.String()))
}