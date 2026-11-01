package service

import (
	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/repository"
)

type Service struct {
	Embedding  *EmbeddingService
	Collection *CollectionService
}

func New(cfg *config.Config, repo *repository.Repository) *Service {
	return &Service{
		Collection: NewCollectionService(repo.Collection),
		Embedding:  NewEmbeddingService(cfg.Embedding, repo.Embedding),
	}
}
