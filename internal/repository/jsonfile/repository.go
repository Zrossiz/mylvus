package jsonfile

import (
	"github.com/Zrossiz/mylvus/internal/config"
)

type JSONRepository struct {
	Embedding *EmbeddingRepository
}

func NewJSONRepository(cfg *config.Config) (*JSONRepository, error) {
	embRepo, err := NewEmbeddingRepository(cfg.Repository.JSONFilePath, cfg.Embedding)
	if err != nil {
		return nil, err
	}

	return &JSONRepository{
		Embedding: embRepo,
	}, nil
}
