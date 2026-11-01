package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/Zrossiz/mylvus/internal/model"
	"github.com/google/uuid"
)

type CollectionRepository struct {
	conn *sql.DB
}

func NewCollectionRepository(conn *sql.DB) *CollectionRepository {
	return &CollectionRepository{
		conn: conn,
	}
}

func (r *CollectionRepository) Create(ctx context.Context, title string, size int64) (uuid.UUID, error) {
	uid := uuid.New()
	_, err := r.conn.ExecContext(ctx, `
		insert into collections (
			id,
			title,
			embedding_size,
			created_at
		) values ($1, $2, $3, $4);
	`, uid, title, size, time.Now())
	if err != nil {
		return uuid.UUID{}, err
	}

	return uid, nil
}

func (r *CollectionRepository) GetAll(ctx context.Context) ([]model.Collection, error) {
	rows, err := r.conn.QueryContext(ctx, `
		select
			id,
			title,
			embedding_size,
			created_at
		from collections;
	`)
	if err != nil {
		return nil, err
	}

	var collections []model.Collection
	for rows.Next() {
		var collection model.Collection

		err = rows.Scan(
			&collection.ID,
			&collection.Title,
			&collection.EmbeddingSize,
			&collection.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		collections = append(collections, collection)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return collections, nil
}
