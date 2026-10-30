package repository

import "github.com/Zrossiz/mylvus/internal/model"

type SQLBinaryRepository struct {
	sqlRepo    SQLMetaRepo
	binaryRepo BinaryEmbeddingRepo
}

func NewSQLBinaryRepository(
	sqlRepo SQLMetaRepo,
	binaryRepo BinaryEmbeddingRepo,
) *SQLBinaryRepository {
	return &SQLBinaryRepository{
		sqlRepo:    sqlRepo,
		binaryRepo: binaryRepo,
	}
}

func (r *SQLBinaryRepository) Create(embedding model.Embedding) error {
	return nil
}

func (r *SQLBinaryRepository) GetAll() []model.Embedding {
	return nil
}
