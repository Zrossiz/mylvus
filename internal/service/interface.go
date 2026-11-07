package service

import (
	"context"

	"github.com/Zrossiz/mylvus/internal/model"
	"github.com/google/uuid"
)

type EmbeddingRepository interface {
	Create(ctx context.Context, collection model.Collection, embedding model.Embedding) error
	GetAll(collectionTitle string) ([]model.Embedding, error)
}

type CollectionRepository interface {
	Create(ctx context.Context, collection model.Collection) (uuid.UUID, error)
	GetByTitle(ctx context.Context, title string) (*model.Collection, error)
	CreateBinaryFile(ctx context.Context, title string) (string, error)
}
