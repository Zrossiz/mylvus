package repository

import (
	"context"

	"github.com/Zrossiz/mylvus/internal/model"
	"github.com/google/uuid"
)

type EmbeddingRepo interface {
	Create(embedding model.Embedding) error
	GetAll() ([]model.Embedding, error)
	// GetAllFromDisk() ([]model.Embedding, error)
}

type CollectionRepo interface {
	Create(ctx context.Context, collection model.Collection) (uuid.UUID, error)
	GetAll(ctx context.Context) ([]model.Collection, error)
}

type ConnectionRepo interface {
	Create(collection, addr string)
	Delete(collection, addr string)
}

type SQLMetaRepo interface {
	Create(meta model.EmbeddingMeta) error
	CreateFreeOffset(offset int64) error
	Delete(id int64, offset int64) error
	GetAll() ([]model.EmbeddingMeta, error)
	GetFreeOffset() (*int64, error)
}

type BinaryEmbeddingRepo interface {
	GetEmbeddingByOffset(filepath string, offset, embLen int64) ([]float32, error)
	InsertAtEnd(filepath string, embedding []float32) (int64, error)
	InsertAtOffset(filepath string, embedding []float32, offset int64) error
}
