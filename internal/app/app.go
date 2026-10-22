package app

import (
	"fmt"
	"net/http"

	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/repository"
	"github.com/Zrossiz/mylvus/internal/service"
	"github.com/Zrossiz/mylvus/internal/transport/rest"
	"github.com/go-chi/chi"
)

func Start() {
	cfg := config.New()

	repo := repository.New(cfg.Repository)
	srv := service.New(repo)

	r := chi.NewRouter()
	rest.New(r, srv)

	address := fmt.Sprintf("%s:%s", cfg.Server.Address, cfg.Server.Port)
	if err := http.ListenAndServe(address, r); err != nil {

	}
}
