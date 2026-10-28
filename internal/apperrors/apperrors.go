package apperrors

import "errors"

var (
	ErrInvalidSearchType   = errors.New("invalid search type")
	ErrInvalidEmbeddingLen = errors.New("invalid embedding len")
)
