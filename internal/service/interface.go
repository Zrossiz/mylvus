package service

import (
	"context"

	"github.com/Zrossiz/mylvus/internal/model"
	"github.com/google/uuid"
)

type EmbeddingRepository interface {
	Create(embedding model.Embedding) error
	GetAll() ([]model.Embedding, error)
}

type CollectionRepository interface {
	Create(ctx context.Context, title string, size int64) (uuid.UUID, error)
	GetAll(ctx context.Context) ([]model.Collection, error)
}
