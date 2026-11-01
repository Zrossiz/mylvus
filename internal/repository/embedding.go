package repository

import (
	"sync"
	"time"

	"github.com/Zrossiz/mylvus/internal/model"
)

type SQLBinaryRepository struct {
	sqlRepo    SQLMetaRepo
	binaryRepo BinaryEmbeddingRepo

	mu   sync.RWMutex
	data []model.Embedding
}

func NewSQLBinaryRepository(
	sqlRepo SQLMetaRepo,
	binaryRepo BinaryEmbeddingRepo,
) (*SQLBinaryRepository, error) {
	repo := SQLBinaryRepository{
		sqlRepo:    sqlRepo,
		binaryRepo: binaryRepo,
	}

	embs, err := repo.GetAllFromDisk()
	if err != nil {
		return nil, err
	}

	repo.data = embs

	return &repo, nil
}

func (r *SQLBinaryRepository) Create(embedding model.Embedding) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	offset, err := r.sqlRepo.GetFreeOffset()
	if err != nil {
		return err
	}

	embeddingVectorBytes := int64(len(embedding.Vector) * 4)

	var newOffset int64
	if offset != nil {
		err = r.binaryRepo.InsertAtOffset(embedding.Vector, *offset)
		if err != nil {
			return err
		}

		newOffset = *offset
	} else {
		endOffset, err := r.binaryRepo.InsertAtEnd(embedding.Vector)
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

	err = r.sqlRepo.Create(meta)
	if err != nil {
		return err
	}

	embs, err := r.GetAllFromDisk()
	if err != nil {
		return err
	}
	r.data = embs

	return nil
}

func (r *SQLBinaryRepository) GetAll() ([]model.Embedding, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	return r.data, nil
}

func (r *SQLBinaryRepository) GetAllFromDisk() ([]model.Embedding, error) {
	embeddingsMeta, err := r.sqlRepo.GetAll()
	if err != nil {
		return nil, err
	}

	var result []model.Embedding
	for _, v := range embeddingsMeta {
		// TODO: убрать хардкод
		curVector, err := r.binaryRepo.GetEmbeddingByOffset(v.Offset, 300)
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

	return result, nil
}
