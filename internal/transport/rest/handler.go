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

func New(r chi.Router, cfg config.Embedding, srv *service.Service) {
	h := Handler{
		cfg:          cfg,
		embeddingSrv: srv.Embedding,
	}

	r.Post("/api/v1/embedding", h.Create)
	r.Post("/api/v1/embedding/closest", h.FindClosest)
	r.Post("/api/v1/embedding/matching", h.FindMatching)
}
