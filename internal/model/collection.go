package model

import (
	"time"

	"github.com/google/uuid"
)

type Collection struct {
	ID            uuid.UUID `json:"id"`
	Title         string    `json:"title"`
	EmbeddingSize int64     `json:"embedding_size"`
	CreatedAt     time.Time `json:"created_at"`
}
