package service

import "github.com/Zrossiz/mylvus/internal/model"

type EmbeddingService struct {
	embeddingRepo EmbeddingRepository
}

func NewEmbeddingService(embeddingRepo EmbeddingRepository) *EmbeddingService {
	return &EmbeddingService{
		embeddingRepo: embeddingRepo,
	}
}

func (s *EmbeddingService) Create(embedding model.Embedding) error {
	return s.embeddingRepo.Create(embedding)
}

func (s *EmbeddingService) GetAll() ([]model.Embedding, error) {
	return s.embeddingRepo.GetAll()
}
