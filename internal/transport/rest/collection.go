package rest

import (
	"encoding/json"
	"net/http"
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

	id, err := h.collectionSrv.Create(ctx, body.Title, body.EmbeddingSize)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(id)
}

func (h *Handler) GetAllCollections(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	collections, err := h.collectionSrv.GetAll(ctx)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(collections)
}
