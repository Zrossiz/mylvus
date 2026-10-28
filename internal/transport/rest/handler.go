package rest

import (
	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/service"
	"github.com/Zrossiz/mylvus/internal/transport"
	"github.com/go-chi/chi"
)

type Handler struct {
	cfg          config.Embedding
	embeddingSrv transport.EmbeddingService
}

func New(r chi.Router, srv *service.Service) {
	h := Handler{}

	r.Post("/api/v1/embedding", h.Create)
	r.Get("/api/v1/embedding/closest", h.FindClosest)
	r.Get("/api/v1/embedding/matchind", h.FindMatching)
}
