package rest

type createEmbeddingDTO struct {
	Value  string    `json:"value"`
	Vector []float32 `json:"vector"`
}

type findEmbeddingsDTO struct {
	ClosestCount int       `json:"closest_count"`
	Threshold    float32   `json:"threshold"`
	Vector       []float32 `json:"vector"`
}

type createCollectionDTO struct {
	Title         string `json:"title"`
	EmbeddingSize int64  `json:"embedding_size"`
}
