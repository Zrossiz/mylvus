package service

import (
	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/repository"
)

type Service struct {
	Embedding *EmbeddingService
}

func New(cfg *config.Config, repo *repository.Repository) *Service {
	return &Service{
		Embedding: NewEmbeddingService(cfg.Embedding, repo.Embedding),
	}
}
