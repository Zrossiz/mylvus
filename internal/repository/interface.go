package repository

import "github.com/Zrossiz/mylvus/internal/model"

type EmbeddingRepo interface {
	Create(embedding model.Embedding) error
	GetAll() []model.Embedding
	GetAllFromDisk() ([]model.Embedding, error)
}
