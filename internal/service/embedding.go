package service

import (
	"time"

	"github.com/Zrossiz/mylvus/internal/apperrors"
	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/model"
	"go.uber.org/zap"
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

func (s *EmbeddingService) FindClosest(
	embedding []float32,
	threshold float32,
	closestCount int,
) ([]model.EmbeddingWithScore, error) {
	switch s.cfg.SearchMethod {
	case string(model.SearchTypeL2):
		return s.findClosestL2(embedding, threshold, closestCount)
	case string(model.SearchTypeCosine):
		start := time.Now()

		embs, err := s.findClosestCosine(embedding, threshold, closestCount)
		if err != nil {
			return nil, err
		}

		diff := time.Since(start)
		zap.L().Info("time to find closest cosine embeddings: ", zap.Int64("service", diff.Microseconds()))

		return embs, nil
	default:
		return nil, apperrors.ErrInvalidSearchType
	}
}

func (s *EmbeddingService) FindMatching(
	embedding []float32,
	threshold float32,
	closestCount int,
) ([]model.EmbeddingWithScore, error) {
	switch s.cfg.SearchMethod {
	case string(model.SearchTypeL2):
		return s.findMatchingL2(embedding, threshold, closestCount)
	case string(model.SearchTypeCosine):
		return s.findMatchingCosine(embedding, threshold, closestCount)
	default:
		return nil, apperrors.ErrInvalidSearchType
	}
}

func (s *EmbeddingService) GetAll() []model.Embedding {
	return s.embeddingRepo.GetAll()
}
