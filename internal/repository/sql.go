package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Zrossiz/mylvus/internal/model"
)

type SQLMetaRepository struct {
	conn *sql.DB
}

func NewSQLMetaRepository(filepath string) (*SQLMetaRepository, error) {
	conn, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, fmt.Errorf("db connect err: %w", err)
	}

	return &SQLMetaRepository{
		conn: conn,
	}, nil
}

func (r *SQLMetaRepository) Create() error {
	return nil
}

func (r *SQLMetaRepository) Delete() error {
	return nil
}

func (r *SQLMetaRepository) CreateFreeOffset() error {
	return nil
}

func (r *SQLMetaRepository) GetFreeOffset() (*int64, error) {
	row := r.conn.QueryRow(`
		DELETE FROM free_offsets
		WHERE offset = (
			SELECT offset
			FROM free_offsets
			ORDER BY offset
			LIMIT 1
		)
		RETURNING offset;
	`)

	var result int64
	if err := row.Scan(&result); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (r *SQLMetaRepository) GetAll() ([]model.EmbeddingMeta, error) {
	rows, err := r.conn.Query(`
		select
			id,
			value,
			offset,
			created_at
		from embeddings_meta;
	`)
	if err != nil {
		return nil, err
	}

	var embeddingsMeta []model.EmbeddingMeta
	for rows.Next() {
		var embeddingMeta model.EmbeddingMeta

		err = rows.Scan(
			&embeddingMeta.ID,
			&embeddingMeta.Value,
			&embeddingMeta.Offset,
			&embeddingMeta.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		embeddingsMeta = append(embeddingsMeta, embeddingMeta)
	}

	if rows.Err() != nil {

	}

	return embeddingsMeta, nil
}
