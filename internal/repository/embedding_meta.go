package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Zrossiz/mylvus/internal/model"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

const nilUUIDValue = "00000000-0000-0000-0000-000000000000"

type SQLMetaRepository struct {
	conn *sql.DB
}

func NewSQLMetaRepository(conn *sql.DB) (*SQLMetaRepository, error) {
	return &SQLMetaRepository{
		conn: conn,
	}, nil
}

func Connect(filepath string) (*sql.DB, error) {
	conn, err := sql.Open("sqlite3", filepath)
	if err != nil {
		return nil, fmt.Errorf("db connect err: %w", err)
	}

	return conn, nil
}

func (r *SQLMetaRepository) Create(meta model.EmbeddingMeta) error {
	if meta.CollectionID.String() == nilUUIDValue {
		return fmt.Errorf("collection id must be a value")
	}

	_, err := r.conn.Exec(`
		insert into embeddings_meta (
			id,
			collection_id,
			value,
			offset,
			created_at
		) values ($1, $2, $3, $4, $5);
	`, uuid.New(), meta.CollectionID, meta.Value, meta.Offset, time.Now())
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLMetaRepository) Delete(id int64, offset int64) error {
	_, err := r.conn.Exec(`delete from embeddings_meta where id = $1`, id)
	if err != nil {
		return err
	}

	err = r.CreateFreeOffset(offset)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLMetaRepository) CreateFreeOffset(offset int64) error {
	_, err := r.conn.Exec(`insert into free_offsets (offset) values ($1);`, offset)
	if err != nil {
		return err
	}

	return nil
}

func (r *SQLMetaRepository) GetFreeOffset() (*int64, error) {
	row := r.conn.QueryRow(`
		delete from free_offsets
		where offset = (
			select offset
			from free_offsets
			order by offset
			limit 1
		)
		returning offset;
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
			collection_id,
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
			&embeddingMeta.CollectionID,
			&embeddingMeta.Value,
			&embeddingMeta.Offset,
			&embeddingMeta.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		embeddingsMeta = append(embeddingsMeta, embeddingMeta)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return embeddingsMeta, nil
}
