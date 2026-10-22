package transport

import "github.com/Zrossiz/mylvus/internal/model"

type EmbeddingService interface {
	Create(embedding model.Embedding) error
	GetAll() ([]model.Embedding, error)
}
