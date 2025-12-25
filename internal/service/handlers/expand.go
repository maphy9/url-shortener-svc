package handlers

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/maphy9/url-shortener-svc/internal/service/errors/apierrors"
	"github.com/maphy9/url-shortener-svc/internal/service/helpers"
	"github.com/maphy9/url-shortener-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
)

func Expand(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)

	request, err := requests.NewExpandRequest(r)
	if err != nil {
		logger.WithError(err).Debug("Invalid request")
		ape.RenderErr(w, apierrors.BadRequest())
		return
	}

	originalUrl, err := helpers.GetOriginalUrl(r, request.Alias)
	if errors.Is(err, sql.ErrNoRows) {
		ape.RenderErr(w, apierrors.NotFound())
		return
	}
	if err != nil {
		logger.WithError(err).Debug("Internal Server Error")
		ape.RenderErr(w, apierrors.InternalServerError())
		return
	}
	ape.Render(w, originalUrl)
}
