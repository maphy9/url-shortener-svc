package requests

import "net/url"

type ShortenURLRequest struct {
	Url string `json:"url"`
}

func (r ShortenURLRequest) IsValid() bool {
	_, err := url.ParseRequestURI(r.Url)
	return err == nil
}
