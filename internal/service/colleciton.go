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

func (s *CollectionService) Create(ctx context.Context, collection model.Collection) (uuid.UUID, error) {
	filepath, err := s.collectionRepo.CreateBinaryFile(ctx, collection.Title)
	if err != nil {
		return uuid.UUID{}, err
	}
	collection.BinaryFilepath = filepath

	return s.collectionRepo.Create(ctx, collection)
}
