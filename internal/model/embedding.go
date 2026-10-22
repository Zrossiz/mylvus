package model

type Embedding struct {
	ID        string       `json:"id"`
	Value     string       `json:"value"`
	Vector    [384]float32 `json:"vector"`
	CreatedAt string       `json:"created_at"`
}
