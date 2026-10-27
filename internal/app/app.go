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
	cfg := config.New()

	log, err := logger.New(cfg.App.LoggerLevel)
	if err != nil {
		return fmt.Errorf("init logger err: %w", err)
	}

	zap.ReplaceGlobals(log)

	repo := repository.New(cfg.Repository)
	srv := service.New(cfg, repo)

	r := chi.NewRouter()
	rest.New(r, srv)

	address := fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port)
	if err := http.ListenAndServe(address, r); err != nil {
		return fmt.Errorf("start http server err: %w", err)
	}

	return nil
}
