package service

import (
	"context"
	"time"

	"github.com/Zrossiz/mylvus/internal/apperrors"
	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/model"
	"go.uber.org/zap"
)

type EmbeddingService struct {
	embeddingRepo  EmbeddingRepository
	collectionRepo CollectionRepository
	cfg            config.Embedding
}

func NewEmbeddingService(
	cfg config.Embedding,
	embeddingRepo EmbeddingRepository,
	collectionRepo CollectionRepository,
) *EmbeddingService {
	return &EmbeddingService{
		cfg:            cfg,
		embeddingRepo:  embeddingRepo,
		collectionRepo: collectionRepo,
	}
}

func (s *EmbeddingService) Create(ctx context.Context, collection string, embedding model.Embedding) error {
	col, err := s.collectionRepo.GetByTitle(ctx, collection)
	if err != nil {
		return err
	}

	embedding.CollectionID = col.ID

	return s.embeddingRepo.Create(ctx, *col, embedding)
}

func (s *EmbeddingService) FindClosest(
	ctx context.Context,
	collectionTitle string,
	embedding []float32,
	threshold float32,
	closestCount int,
) ([]model.EmbeddingWithScore, error) {
	switch s.cfg.SearchMethod {
	case string(model.SearchTypeL2):
		return s.findClosestL2(collectionTitle, embedding, threshold, closestCount)
	case string(model.SearchTypeCosine):
		start := time.Now()

		embs, err := s.findClosestCosine(collectionTitle, embedding, threshold, closestCount)
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
	ctx context.Context,
	collectionTitle string,
	embedding []float32,
	threshold float32,
	closestCount int,
) ([]model.EmbeddingWithScore, error) {
	switch s.cfg.SearchMethod {
	case string(model.SearchTypeL2):
		return s.findMatchingL2(collectionTitle, embedding, threshold, closestCount)
	case string(model.SearchTypeCosine):
		return s.findMatchingCosine(collectionTitle, embedding, threshold, closestCount)
	default:
		return nil, apperrors.ErrInvalidSearchType
	}
}

func (s *EmbeddingService) GetAll(collectionTitle string) ([]model.Embedding, error) {
	return s.embeddingRepo.GetAll(collectionTitle)
}
