package repository

import (
	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/repository/jsonfile"
)

type Repository struct {
	Embedding EmbeddingRepo
}

func New(cfg *config.Config) (*Repository, error) {
	var repo Repository

	if cfg.Repository.JSONFilePath != "" {
		jsonRepo, err := jsonfile.NewJSONRepository(cfg)
		if err != nil {
			return nil, err
		}

		repo.Embedding = jsonRepo.Embedding
	}

	return &repo, nil
}
