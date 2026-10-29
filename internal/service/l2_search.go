package service

import (
	"fmt"
	"math"
	"sort"

	"github.com/Zrossiz/mylvus/internal/model"
)

func (s *EmbeddingService) findClosestL2(
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

		score, err := s.getMatchingScoreL2(embedding, curEmb)
		if err != nil {
			return nil, err
		}

		if score < float64(threshold) {
			matching = append(matching, model.EmbeddingWithScore{
				Embedding: curEmb,
				Score:     score,
			})
		}
	}

	// from min to max
	sort.Slice(matching, func(i, j int) bool {
		return matching[i].Score < matching[j].Score
	})

	var result []model.EmbeddingWithScore
	for i := 0; i < len(matching) && closestCount > 0; i++ {
		result = append(result, matching[i])
		closestCount--
	}

	return result, nil
}

func (s *EmbeddingService) findMatchingL2(
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

		score, err := s.getMatchingScoreL2(embedding, curEmb)
		if err != nil {
			return nil, err
		}

		if score < float64(threshold) {
			result = append(result, model.EmbeddingWithScore{
				Embedding: curEmb,
				Score:     score,
			})
			closestCount--
		}
	}

	return result, nil
}

func (s *EmbeddingService) getMatchingScoreL2(
	query []float32,
	reference model.Embedding,
) (float64, error) {
	if len(query) != len(reference.Vector) {
		return 0, fmt.Errorf("len of vectors must be equal")
	}

	var sum float32

	for i := 0; i < len(query); i++ {
		subtraction := reference.Vector[i] - query[i]

		sum += subtraction * subtraction
	}

	d := math.Sqrt(float64(sum))

	return d, nil
}
