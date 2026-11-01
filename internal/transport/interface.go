package transport

import (
	"context"

	"github.com/Zrossiz/mylvus/internal/model"
	"github.com/google/uuid"
)

type EmbeddingService interface {
	Create(embedding model.Embedding) error
	FindClosest(embedding []float32, threshold float32, closestCount int) ([]model.EmbeddingWithScore, error)
	FindMatching(embedding []float32, threshold float32, closestCount int) ([]model.EmbeddingWithScore, error)
}

type CollectionService interface {
	Create(ctx context.Context, title string, size int64) (uuid.UUID, error)
	GetAll(ctx context.Context) ([]model.Collection, error)
}
