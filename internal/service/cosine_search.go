package service

import (
	"fmt"
	"math"
	"sort"

	"github.com/Zrossiz/mylvus/internal/model"
)

func (s *EmbeddingService) findClosestCosine(
	embedding []float32,
	threshold float32,
	closestCount int,
) ([]model.EmbeddingWithScore, error) {
	embs, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	var matching []model.EmbeddingWithScore
	for i := 0; i < len(embs) && closestCount > 0; i++ {
		curEmb := embs[i]

		score, err := s.getMatchingScoreCosine(embedding, curEmb)
		if err != nil {
			return nil, err
		}

		if score > float64(threshold) {
			matching = append(matching, model.EmbeddingWithScore{
				Embedding: curEmb,
				Score:     score,
			})
		}
	}

	// from max to min
	sort.Slice(matching, func(i, j int) bool {
		return matching[i].Score > matching[j].Score
	})

	var result []model.EmbeddingWithScore
	for i := 0; i < len(matching) && closestCount > 0; i++ {
		result = append(result, matching[i])
		closestCount--
	}

	return result, nil
}

func (s *EmbeddingService) findMatchingCosine(
	embedding []float32,
	threshold float32,
	closestCount int,
) ([]model.EmbeddingWithScore, error) {
	embs, err := s.GetAll()
	if err != nil {
		return nil, err
	}

	var result []model.EmbeddingWithScore
	for i := 0; i < len(embs) && closestCount > 0; i++ {
		curEmb := embs[i]

		score, err := s.getMatchingScoreCosine(embedding, curEmb)
		if err != nil {
			return nil, err
		}

		if score > float64(threshold) {
			result = append(result, model.EmbeddingWithScore{
				Embedding: curEmb,
				Score:     score,
			})
			closestCount--
		}
	}

	return result, nil
}

func (s *EmbeddingService) getMatchingScoreCosine(
	query []float32,
	reference model.Embedding,
) (float64, error) {
	if len(query) != len(reference.Vector) {
		return 0, fmt.Errorf("len of vectors must be equal")
	}

	var sumOfCoordinates float32
	var sumQuery float32
	var sumRef float32

	for i := 0; i < len(query); i++ {
		sumOfCoordinates += query[i] * reference.Vector[i]
		sumQuery += query[i] * query[i]
		sumRef += reference.Vector[i] * reference.Vector[i]
	}

	normQuery := math.Sqrt(float64(sumQuery))
	normRef := math.Sqrt(float64(sumRef))

	denominator := float32(normQuery) * float32(normRef)
	cosine := sumOfCoordinates / denominator

	return float64(cosine), nil
}
