package rest

type createEmbeddingDTO struct {
	Value  string    `json:"value"`
	Vector []float32 `json:"vector"`
}

type findEmbeddingsDTO struct {
	SearchMethod string    `json:"search_method"`
	Vector       []float32 `json:"vector"`
}
