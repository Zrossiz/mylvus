package rest

import (
	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/service"
	"github.com/Zrossiz/mylvus/internal/transport"
	"github.com/go-chi/chi"
)

type Handler struct {
	cfg           config.Embedding
	embeddingSrv  transport.EmbeddingService
	collectionSrv transport.CollectionService
}

func New(r chi.Router, cfg config.Embedding, srv *service.Service) {
	h := Handler{
		cfg:           cfg,
		embeddingSrv:  srv.Embedding,
		collectionSrv: srv.Collection,
	}

	r.Post("/api/v1/collections", h.CreateCollection)
	r.Post("/api/v1/collections/{title}/embeddings", h.CreateEmbedding)
	r.Post("/api/v1/collections/{title}/embeddings/closest", h.FindClosest)
	r.Post("/api/v1/collections/{title}/embeddings/matching", h.FindMatching)
}
