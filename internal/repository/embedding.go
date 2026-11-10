package repository

import (
	"context"
	"database/sql"
	"errors"
	"sync"
	"time"

	"github.com/Zrossiz/mylvus/internal/model"
)

type SQLBinaryRepository struct {
	sqlMetaRepo       SQLMetaRepo
	sqlCollectionRepo SQLCollectionRepo
	binaryRepo        BinaryEmbeddingRepo
	conn              *sql.DB

	mu   sync.RWMutex
	data map[string][]model.Embedding
}

func NewSQLBinaryRepository(
	sqlMetaRepo SQLMetaRepo,
	binaryRepo BinaryEmbeddingRepo,
	sqlCollectionRepo SQLCollectionRepo,
	conn *sql.DB,
) (*SQLBinaryRepository, error) {
	repo := SQLBinaryRepository{
		sqlMetaRepo:       sqlMetaRepo,
		binaryRepo:        binaryRepo,
		sqlCollectionRepo: sqlCollectionRepo,
		conn:              conn,
		data:              make(map[string][]model.Embedding),
	}

	return &repo, nil
}

func (r *SQLBinaryRepository) Create(ctx context.Context, collection model.Collection, embedding model.Embedding) error {
	tx, err := r.conn.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback()

	offset, err := r.sqlMetaRepo.GetFreeOffset(ctx, tx, collection.ID)
	if err != nil {
		return err
	}

	embeddingVectorBytes := int64(len(embedding.Vector) * 4)

	var newOffset int64
	if offset != nil {
		err = r.binaryRepo.InsertAtOffset(collection.BinaryFilepath, embedding.Vector, *offset)
		if err != nil {
			return err
		}

		newOffset = *offset
	} else {
		endOffset, err := r.binaryRepo.InsertAtEnd(collection.BinaryFilepath, embedding.Vector)
		if err != nil {
			return err
		}

		newOffset = endOffset - embeddingVectorBytes
	}

	meta := model.EmbeddingMeta{
		CollectionID: embedding.CollectionID,
		Value:        embedding.Value,
		Offset:       newOffset,
		CreatedAt:    time.Now(),
	}

	err = r.sqlMetaRepo.Create(ctx, tx, meta)
	if err != nil {
		if offset == nil {
			truncErr := r.binaryRepo.Truncate(collection.BinaryFilepath, newOffset)
			return errors.Join(err, truncErr)
		}

		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	embs, err := r.GetAllFromDisk(ctx, collection)
	if err != nil {
		return err
	}
	r.data[collection.Title] = embs

	return nil
}

func (r *SQLBinaryRepository) GetAll(ctx context.Context, collectionTitle string) ([]model.Embedding, error) {
	r.mu.RLock()
	embs, _ := r.data[collectionTitle]
	r.mu.RUnlock()

	if len(embs) != 0 {
		return embs, nil
	}

	diskEmbs, err := r.GetAllFromDiskByTitle(ctx, collectionTitle)
	if err != nil {
		return nil, err
	}

	return diskEmbs, nil
}

func (r *SQLBinaryRepository) GetAllFromDiskByTitle(ctx context.Context, collectionTitle string) ([]model.Embedding, error) {
	collection, err := r.sqlCollectionRepo.GetByTitle(ctx, collectionTitle)
	if err != nil {
		return nil, err
	}

	embs, err := r.GetAllFromDisk(ctx, *collection)
	if err != nil {
		return nil, err
	}

	return embs, nil
}

func (r *SQLBinaryRepository) GetAllFromDisk(ctx context.Context, collection model.Collection) ([]model.Embedding, error) {
	embeddingsMeta, err := r.sqlMetaRepo.GetAll(ctx, collection.ID)
	if err != nil {
		return nil, err
	}

	var result []model.Embedding
	for _, v := range embeddingsMeta {
		curVector, err := r.binaryRepo.GetEmbeddingByOffset(collection.BinaryFilepath, v.Offset, collection.EmbeddingSize)
		if err != nil {
			return nil, err
		}

		result = append(result, model.Embedding{
			ID:           v.ID,
			CollectionID: v.CollectionID,
			Value:        v.Value,
			Vector:       curVector,
			CreatedAt:    v.CreatedAt,
		})
	}

	r.mu.Lock()
	r.data[collection.Title] = result
	r.mu.Unlock()

	return result, nil
}
