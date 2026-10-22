package jsonfile

import (
	"github.com/Zrossiz/mylvus/internal/config"
)

type JSONRepository struct {
	Embedding *EmbeddingRepository
}

func NewJSONRepository(cfg config.Repository) *JSONRepository {
	return &JSONRepository{
		Embedding: NewEmbeddingRepository(cfg.JSONFilePath),
	}
}
