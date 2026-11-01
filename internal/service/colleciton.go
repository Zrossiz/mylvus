package service

import (
	"context"

	"github.com/Zrossiz/mylvus/internal/model"
	"github.com/google/uuid"
)

type CollectionService struct {
	collectionRepo CollectionRepository
}

func NewCollectionService(collectionRepo CollectionRepository) *CollectionService {
	return &CollectionService{
		collectionRepo: collectionRepo,
	}
}

func (s *CollectionService) Create(ctx context.Context, title string, size int64) (uuid.UUID, error) {
	return s.collectionRepo.Create(ctx, title, size)
}

func (s *CollectionService) GetAll(ctx context.Context) ([]model.Collection, error) {
	return s.collectionRepo.GetAll(ctx)
}
