package service

import (
	"fmt"
	"math"

	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/model"
)

type EmbeddingService struct {
	embeddingRepo EmbeddingRepository
	cfg           config.Embedding
}

func NewEmbeddingService(cfg config.Embedding, embeddingRepo EmbeddingRepository) *EmbeddingService {
	return &EmbeddingService{
		cfg:           cfg,
		embeddingRepo: embeddingRepo,
	}
}

func (s *EmbeddingService) Create(embedding model.Embedding) error {
	return s.embeddingRepo.Create(embedding)
}

func (s *EmbeddingService) GetClosest(
	embedding []float32,
	threshold float32,
	closestCount int,
) ([]model.Embedding, error) {
	embs, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	var result []model.Embedding
	for i := 0; i < len(embs) && closestCount > 0; i++ {
		curEmb := embs[i]

		isFound, err := s.isMatching(embedding, curEmb, threshold)
		if err != nil {
			return nil, err
		}

		if isFound {
			result = append(result, curEmb)
			closestCount--
		}
	}

	return result, nil
}

func (s *EmbeddingService) GetAll() ([]model.Embedding, error) {
	return s.embeddingRepo.GetAll()
}

func (s *EmbeddingService) isMatching(
	query []float32,
	reference model.Embedding,
	threshold float32,
) (bool, error) {
	if len(query) != len(reference.Vector) {
		return false, fmt.Errorf("len of vectors must be equal")
	}

	var sumOfCoordinates float32
	var sumQuery float32
	var sumRef float32

	for i := 0; i < len(query); i++ {
		sumOfCoordinates += query[i] * reference.Vector[i]
		sumQuery += query[i] * query[i]
		sumRef += reference.Vector[i] * reference.Vector[i]
	}

	normQuery := math.Sqrt(float64(sumQuery))
	normRef := math.Sqrt(float64(sumRef))

	denominator := float32(normQuery) * float32(normRef)
	cosine := sumOfCoordinates / denominator

	if cosine > threshold {
		return true, nil
	}

	return false, nil
}
