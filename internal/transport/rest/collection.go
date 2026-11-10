package rest

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/Zrossiz/mylvus/internal/model"
	"github.com/google/uuid"
)

func (h *Handler) CreateCollection(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	var body createCollectionDTO

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if strings.Contains(body.Title, "../") || body.Title == "" {
		http.Error(rw, "invalid collection title", http.StatusBadRequest)
		return
	}

	collection := model.Collection{
		ID:            uuid.New(),
		Title:         body.Title,
		EmbeddingSize: body.EmbeddingSize,
	}

	id, err := h.collectionSrv.Create(ctx, collection)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(id)
}
