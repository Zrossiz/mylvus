package app

import (
	"fmt"
	"net/http"

	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/repository"
	"github.com/Zrossiz/mylvus/internal/service"
	"github.com/Zrossiz/mylvus/internal/transport/rest"
	"github.com/Zrossiz/mylvus/pkg/logger"
	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func Start() error {
	cfg, err := config.New()
	if err != nil {
		return err
	}

	log, err := logger.New(cfg.App.LoggerLevel)
	if err != nil {
		return fmt.Errorf("init logger err: %w", err)
	}

	zap.ReplaceGlobals(log)

	repo, err := repository.New(cfg)
	if err != nil {
		return err
	}
	srv := service.New(cfg, repo)

	r := chi.NewRouter()
	rest.New(r, cfg.Embedding, srv)

	zap.L().Sugar().Infof("starting server on %s:%s", cfg.Server.Address, cfg.Server.Port)
	address := fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port)
	if err := http.ListenAndServe(address, r); err != nil {
		return fmt.Errorf("start http server err: %w", err)
	}

	return nil
}
