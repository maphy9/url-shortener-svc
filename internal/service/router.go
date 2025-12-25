package service

import (
	"github.com/go-chi/chi"
	"github.com/maphy9/url-shortener-svc/internal/config"
	"github.com/maphy9/url-shortener-svc/internal/data/pg"
	"github.com/maphy9/url-shortener-svc/internal/service/handlers"
	"gitlab.com/distributed_lab/ape"
)

func (s *service) router(cfg config.Config) chi.Router {
	r := chi.NewRouter()

	r.Use(
		ape.RecoverMiddleware(s.log),
		ape.LoganMiddleware(s.log),
		ape.CtxMiddleware(
			handlers.CtxLog(s.log),
			handlers.CtxDB(pg.NewMasterQ(cfg.DB())),
		),
	)
	r.Route("/", func(r chi.Router) {
		r.Post("/", handlers.ShortenURL)
		r.Get("/{alias}", handlers.ExpandURL)
	})

	return r
}
