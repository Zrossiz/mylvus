package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/repository"
	"github.com/Zrossiz/mylvus/internal/service"
	"github.com/Zrossiz/mylvus/internal/transport/rest"
	"github.com/Zrossiz/mylvus/pkg/logger"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

type App struct {
	cfg        *config.Config
	repo       *repository.Repository
	srv        *service.Service
	ctx        context.Context
	r          *chi.Mux
	httpServer *http.Server
	errCh      chan error
}

func New() (*App, error) {
	var a App

	cfg, err := config.New()
	if err != nil {
		return nil, err
	}
	a.cfg = cfg

	log, err := logger.New(cfg.App.LoggerLevel)
	if err != nil {
		return nil, fmt.Errorf("init logger err: %w", err)
	}

	zap.ReplaceGlobals(log)

	repo, err := repository.New(cfg)
	if err != nil {
		return nil, err
	}
	a.repo = repo

	srv := service.New(cfg, repo)
	a.srv = srv

	r := chi.NewRouter()
	rest.New(r, cfg.Embedding, srv)
	a.r = r

	address := fmt.Sprintf("%s:%s", a.cfg.Server.Address, a.cfg.Server.Port)
	a.httpServer = &http.Server{
		Handler: a.r,
		Addr:    address,
	}

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	a.ctx = ctx

	a.errCh = make(chan error)

	return &a, nil
}

func (a *App) Start() error {
	go func() error {
		zap.L().Sugar().Infof("starting server on %s", a.httpServer.Addr)
		if err := a.httpServer.ListenAndServe(); err != nil {
			a.errCh <- fmt.Errorf("start http server err: %w", err)
		}

		return nil
	}()

	return nil
}

func (a *App) Wait() error {
	select {
	case err := <-a.errCh:
		zap.L().Error(err.Error())
	case <-a.ctx.Done():
	}

	err := a.Stop()
	if err != nil {
		return err
	}

	return nil
}

func (a *App) Stop() error {
	zap.L().Info("app stopping...")
	err := a.httpServer.Close()
	if err != nil {
		return err
	}

	err = a.repo.Close()
	if err != nil {
		return err
	}

	zap.L().Info("app successfully stopped")

	return nil
}
