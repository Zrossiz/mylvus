package repository

import "github.com/Zrossiz/mylvus/internal/model"

type EmbeddingRepo interface {
	Create(embedding model.Embedding) error
	GetAll() ([]model.Embedding, error)
	// GetAllFromDisk() ([]model.Embedding, error)
}

type SQLMetaRepo interface {
	Create(meta model.EmbeddingMeta) error
	CreateFreeOffset(offset int64) error
	Delete(id int64, offset int64) error
	GetAll() ([]model.EmbeddingMeta, error)
	GetFreeOffset() (*int64, error)
}

type BinaryEmbeddingRepo interface {
	GetEmbeddingByOffset(offset int64) ([]float32, error)
	InsertAtEnd(embedding []float32) (int64, error)
	InsertAtOffset(embedding []float32, offset int64) error
}
