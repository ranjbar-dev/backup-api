package api

import (
	"context"

	"github.com/ranjbar-dev/backup-api/internal/config"
	"github.com/ranjbar-dev/backup-api/internal/httpserver"
)

type Api struct {
	ctx context.Context
	hs  *httpserver.HttpServer
}

func (a *Api) Start() error {

	a.registerRoutes()

	return a.hs.Serve()
}

func NewApi(ctx context.Context) *Api {

	return &Api{
		ctx: ctx,
		hs:  httpserver.NewHttpServer(config.Single.ApiHost(), config.Single.ApiPort(), config.Single.ApiSslCrtLocation(), config.Single.ApiSslKeyLocation(), config.Single.ApiDebug()),
	}
}
