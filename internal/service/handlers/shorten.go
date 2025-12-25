package handlers

import (
	"net/http"

	"github.com/maphy9/url-shortener-svc/internal/service/errors/apierrors"
	"github.com/maphy9/url-shortener-svc/internal/service/helpers"
	"github.com/maphy9/url-shortener-svc/internal/service/requests"
	"gitlab.com/distributed_lab/ape"
)

func Shorten(w http.ResponseWriter, r *http.Request) {
	logger := helpers.Log(r)
	request, err := requests.NewShortenRequest(r)
	if err != nil {
		logger.WithError(err).Debug("Invalid Request")
		ape.RenderErr(w, apierrors.BadRequest())
		return
	}

	shortUrl, err := helpers.GetShortUrl(r, request.Url)
	if err != nil {
		logger.WithError(err).Debug("Internal Server Error")
		ape.RenderErr(w, apierrors.InternalServerError())
		return
	}
	ape.Render(w, shortUrl)
}
