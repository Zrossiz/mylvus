package transport

import (
	"context"

	"github.com/Zrossiz/mylvus/internal/model"
	"github.com/google/uuid"
)

type EmbeddingService interface {
	Create(ctx context.Context, collection string, embedding model.Embedding) error
	FindClosest(ctx context.Context, collectionTitle string, embedding []float32, threshold float32, closestCount int) ([]model.EmbeddingWithScore, error)
	FindMatching(ctx context.Context, collectionTitle string, embedding []float32, threshold float32, closestCount int) ([]model.EmbeddingWithScore, error)
	GetAll(ctx context.Context, collectionTitle string) ([]model.Embedding, error)
}

type CollectionService interface {
	Create(ctx context.Context, collection model.Collection) (uuid.UUID, error)
}
