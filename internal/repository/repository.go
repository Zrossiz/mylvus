package repository

import "github.com/Zrossiz/mylvus/internal/config"

type Repository struct {
	Embedding EmbeddingRepo
}

func New(cfg *config.Config) (*Repository, error) {
	var repo Repository

	if cfg.Repository.JSONFilePath != "" {
		jsonRepo, err := NewJSONRepository(cfg.Repository.JSONFilePath, cfg.Embedding)
		if err != nil {
			return nil, err
		}

		repo.Embedding = jsonRepo
	} else if cfg.Repository.BinaryFilePath != "" && cfg.Repository.SQLFilePath != "" {
		sqlRepo, err := NewSQLMetaRepository(cfg.Repository.SQLFilePath)
		if err != nil {
			return nil, err
		}

		binaryRepo := NewBinaryEmbeddingRepository(cfg.Repository.BinaryFilePath, cfg.Embedding.Length)
		combinedRepo := NewSQLBinaryRepository(sqlRepo, binaryRepo)

		repo.Embedding = combinedRepo
	}

	return &repo, nil
}
