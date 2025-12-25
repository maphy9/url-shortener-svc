package requests

import (
	"encoding/json"
	"net/http"
	"net/url"
)

func NewShortenRequest(r *http.Request) (ShortenRequest, error) {
	var request ShortenRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		return request, err
	}
	return request, request.validate()
}

type ShortenRequest struct {
	Url string `json:"url"`
}

func (r ShortenRequest) validate() error {
	_, err := url.ParseRequestURI(r.Url)
	return err
}
