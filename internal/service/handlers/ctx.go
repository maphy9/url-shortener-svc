package handlers

import (
	"context"
	"net/http"

	"github.com/maphy9/url-shortener-svc/internal/service/data"
	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	aliasManagerCtxKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func CtxAliasManager(aliasManager data.AliasManager) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, aliasManagerCtxKey, aliasManager)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func AliasManager(r *http.Request) data.AliasManager {
	return r.Context().Value(aliasManagerCtxKey).(data.AliasManager)
}
