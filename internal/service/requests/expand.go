package requests

import (
	"errors"
	"net/http"
	"regexp"

	"github.com/go-chi/chi"
)

func NewExpandRequest(r *http.Request) (ExpandRequest, error) {
	alias := chi.URLParam(r, "alias")
	request := ExpandRequest{alias}
	return request, request.validate()
}

type ExpandRequest struct {
	Alias string `json:"alias"`
}

func (r ExpandRequest) validate() error {
	matched, err := regexp.MatchString(`^[0-9a-zA-Z]{1,7}$`, r.Alias)
	if err != nil {
		return err
	}
	if !matched {
		return errors.New("Invalid alias format")
	}
	return nil
}