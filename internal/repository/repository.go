package repository

import (
	"fmt"

	"github.com/Zrossiz/mylvus/internal/config"
	"go.uber.org/zap"
)

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
		zap.L().Info("app inited with JSON db")
	} else if cfg.Repository.BinaryFilePath != "" && cfg.Repository.SQLFilePath != "" {
		sqlRepo, err := NewSQLMetaRepository(cfg.Repository.SQLFilePath)
		if err != nil {
			return nil, err
		}

		binaryRepo := NewBinaryEmbeddingRepository(cfg.Repository.BinaryFilePath, cfg.Embedding.Length)
		combinedRepo, err := NewSQLBinaryRepository(sqlRepo, binaryRepo)
		if err != nil {
			return nil, err
		}

		repo.Embedding = combinedRepo
		zap.L().Info("app inited with combined db")
	} else {
		return nil, fmt.Errorf("invalid db format")
	}

	return &repo, nil
}
