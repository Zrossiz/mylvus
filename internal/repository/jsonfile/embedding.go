package jsonfile

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Zrossiz/mylvus/internal/apperrors"
	"github.com/Zrossiz/mylvus/internal/config"
	"github.com/Zrossiz/mylvus/internal/model"
)

type EmbeddingRepository struct {
	cfg      config.Embedding
	filepath string
}

func NewEmbeddingRepository(filepath string, cfg config.Embedding) *EmbeddingRepository {
	return &EmbeddingRepository{
		filepath: filepath,
		cfg:      cfg,
	}
}

func (r *EmbeddingRepository) Create(embedding model.Embedding) error {
	if r.cfg.Length != len(embedding.Vector) {
		return apperrors.ErrInvalidEmbeddingLen
	}

	file, err := os.OpenFile(r.filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return fmt.Errorf("open json file err: %w", err)
	}
	defer file.Close()

	embeddings, err := r.getAllByFile(file)
	if err != nil {
		return fmt.Errorf("get all embeddings err: %w", err)
	}

	embeddings = append(embeddings, embedding)

	file.Seek(0, 0)
	file.Truncate(0)
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	encoder.Encode(embeddings)

	return nil
}

func (r *EmbeddingRepository) GetAll() ([]model.Embedding, error) {
	file, err := os.OpenFile(r.filepath, os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("open json file err: %w", err)
	}
	defer file.Close()

	return r.getAllByFile(file)
}

func (r *EmbeddingRepository) getAllByFile(file *os.File) ([]model.Embedding, error) {
	var embeddings []model.Embedding
	decoder := json.NewDecoder(file)
	err := decoder.Decode(&embeddings)
	if err != nil {
		return nil, fmt.Errorf("decode file embeddings err: %w", err)
	}

	return embeddings, nil
}
