package service

import "github.com/Zrossiz/mylvus/internal/repository"

type Service struct {
	Embedding *EmbeddingService
}

func New(repo *repository.Repository) *Service {
	return &Service{
		Embedding: NewEmbeddingService(repo.Embedding),
	}
}
