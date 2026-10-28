package transport

import "github.com/Zrossiz/mylvus/internal/model"

type EmbeddingService interface {
	Create(embedding model.Embedding) error
	FindClosest(embedding []float32, threshold float32, closestCount int) ([]model.Embedding, error)
	FindMatching(embedding []float32, threshold float32, closestCount int) ([]model.Embedding, error)
}
