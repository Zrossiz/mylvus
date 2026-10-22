package service

import (
	"github.com/Zrossiz/mylvus/internal/model"
)

type EmbeddingRepository interface {
	Create(embedding model.Embedding) error
	GetAll() ([]model.Embedding, error)
}
