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

type EmbeddingMeta struct {
	ID        int       `json:"id"`
	Value     string    `json:"value"`
	Offset    int       `json:"offset"`
	CreatedAt time.Time `json:"created_at"`
}

type EmbeddingWithScore struct {
	Embedding Embedding `json:"embedding"`
	Score     float64   `json:"score"`
}
