package repository

import (
	"database/sql"

	"github.com/Zrossiz/mylvus/internal/config"
)

type Repository struct {
	Embedding  EmbeddingRepo
	Collection SQLCollectionRepo

	conn *sql.DB
}

func New(cfg *config.Config) (*Repository, error) {
	conn, err := Connect(cfg.Repository.SQLFilePath)
	if err != nil {
		return nil, err
	}

	var repo Repository

	sqlMetaRepo, err := NewSQLMetaRepository(conn)
	if err != nil {
		return nil, err
	}

	binaryRepo := NewBinaryEmbeddingRepository()
	repo.Collection = NewCollectionRepository(conn)
	combinedRepo, err := NewSQLBinaryRepository(sqlMetaRepo, binaryRepo, repo.Collection, conn)
	if err != nil {
		return nil, err
	}

	repo.Embedding = combinedRepo

	return &repo, nil
}

func (r *Repository) Close() error {
	if r.conn != nil {
		return r.conn.Close()
	}

	return nil
}
