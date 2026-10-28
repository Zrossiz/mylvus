package repository

import (
	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/repository/jsonfile"
)

type Repository struct {
	Embedding EmbeddingRepo
}

func New(cfg *config.Config) *Repository {
	var repo Repository

	if cfg.Repository.JSONFilePath != "" {
		jsonRepo := jsonfile.NewJSONRepository(cfg)

		repo.Embedding = jsonRepo.Embedding
	}

	return &repo
}
