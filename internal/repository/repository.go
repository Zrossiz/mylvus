package repository

import (
	"database/sql"

	"github.com/Zrossiz/mylvus/internal/config"
)

type Repository struct {
	Embedding  EmbeddingRepo
	Collection CollectionRepo

	conn *sql.DB
}

func New(cfg *config.Config) (*Repository, error) {
	conn, err := Connect(cfg.Repository.SQLFilePath)
	if err != nil {
		return nil, err
	}

	var repo Repository

	sqlRepo, err := NewSQLMetaRepository(conn)
	if err != nil {
		return nil, err
	}

	binaryRepo := NewBinaryEmbeddingRepository()
	combinedRepo, err := NewSQLBinaryRepository(sqlRepo, binaryRepo)
	if err != nil {
		return nil, err
	}

	repo.Embedding = combinedRepo
	repo.Collection = NewCollectionRepository(conn)

	return &repo, nil
}

func (r *Repository) Close() error {
	if r.conn != nil {
		return r.conn.Close()
	}

	return nil
}
