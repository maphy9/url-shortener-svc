package requests

import "net/url"

type ShortenURLRequest struct {
	URL string `json:"url"`
}

func (r ShortenURLRequest) IsValid() bool {
	_, err := url.ParseRequestURI(r.URL)
	return err == nil
}
