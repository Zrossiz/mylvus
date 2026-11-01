package rest

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/Zrossiz/mylvus/internal/model"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func (h *Handler) CreateEmbedding(rw http.ResponseWriter, r *http.Request) {
	var body createEmbeddingDTO

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	parsedUUID, err := uuid.Parse(body.CollectionID)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	emb := model.Embedding{
		ID:           uuid.New(),
		CollectionID: parsedUUID,
		Value:        body.Value,
		Vector:       body.Vector,
		CreatedAt:    time.Now(),
	}

	err = h.embeddingSrv.Create(emb)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (h *Handler) FindMatching(rw http.ResponseWriter, r *http.Request) {
	var body findEmbeddingsDTO

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	embs, err := h.embeddingSrv.FindMatching(body.Vector, body.Threshold, body.ClosestCount)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(embs)

}

func (h *Handler) FindClosest(rw http.ResponseWriter, r *http.Request) {
	var body findEmbeddingsDTO
	start := time.Now()

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	embs, err := h.embeddingSrv.FindClosest(body.Vector, body.Threshold, body.ClosestCount)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	diff := time.Since(start)
	zap.L().Info("time to find closest transport cosine embeddings: ", zap.Int64("transport", diff.Microseconds()))

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(embs)
}
