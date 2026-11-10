package repository

import (
	"context"
	"database/sql"

	"github.com/Zrossiz/mylvus/internal/model"
	"github.com/google/uuid"
)

type EmbeddingRepo interface {
	Create(ctx context.Context, collection model.Collection, embedding model.Embedding) error
	GetAll(ctx context.Context, collectionTitle string) ([]model.Embedding, error)
	GetAllFromDisk(ctx context.Context, collectiont model.Collection) ([]model.Embedding, error)
}

type SQLCollectionRepo interface {
	Create(ctx context.Context, collection model.Collection) (uuid.UUID, error)
	GetByTitle(ctx context.Context, title string) (*model.Collection, error)
	CreateBinaryFile(ctx context.Context, title string) (string, error)
}

type SQLMetaRepo interface {
	Create(ctx context.Context, tx *sql.Tx, meta model.EmbeddingMeta) error
	CreateFreeOffset(ctx context.Context, collectionID uuid.UUID, offset int64) error
	Delete(ctx context.Context, collectionID, embeddingID uuid.UUID, offset int64) error
	GetAll(ctx context.Context, collectionID uuid.UUID) ([]model.EmbeddingMeta, error)
	GetFreeOffset(ctx context.Context, tx *sql.Tx, collectionID uuid.UUID) (*int64, error)
}

type BinaryEmbeddingRepo interface {
	GetEmbeddingByOffset(filepath string, offset, embLen int64) ([]float32, error)
	InsertAtEnd(filepath string, embedding []float32) (int64, error)
	InsertAtOffset(filepath string, embedding []float32, offset int64) error
	Truncate(filepath string, offset int64) error
}
