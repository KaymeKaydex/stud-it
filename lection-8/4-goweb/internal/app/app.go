package app

import (
	"context"

	"test/internal/pkg/api"
	"test/internal/pkg/service"
)

type App struct {
	ctx context.Context

	srv *service.Service
}

func New(ctx context.Context) (*App, error) {
	srv, err := service.New(ctx)
	if err != nil {
		return nil, err
	}

	return &App{
		ctx: ctx,
		srv: srv,
	}, nil
}

func (a *App) Run() error {
	router, err := api.New(a.ctx, a.srv)
	if err != nil {
		return err
	}

	return router.Run()
}
