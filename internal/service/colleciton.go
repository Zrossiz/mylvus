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
	collectionID, err := s.collectionRepo.Create(ctx, collection)
	if err != nil {
		return uuid.UUID{}, err
	}

	filepath, err := s.collectionRepo.CreateBinaryFile(ctx, collection.Title)
	if err != nil {
		return uuid.UUID{}, err
	}
	collection.BinaryFilepath = filepath

	return collectionID, nil
}
