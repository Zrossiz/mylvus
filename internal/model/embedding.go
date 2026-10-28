package model

import (
	"time"

	"github.com/google/uuid"
)

type SearchType string

const (
	SearchTypeL2     SearchType = "L2"
	SearchTypeCosine SearchType = "cosine"
)

type Embedding struct {
	ID        uuid.UUID `json:"id"`
	Value     string    `json:"value"`
	Vector    []float32 `json:"vector"`
	CreatedAt time.Time `json:"created_at"`
}

type EmbeddingWithScore struct {
	Embedding Embedding
	Score     float64
}
