package repository

import (
	"context"
	"database/sql"
	"fmt"
	"os"
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

func (r *CollectionRepository) Create(ctx context.Context, collection model.Collection) (uuid.UUID, error) {
	_, err := r.conn.ExecContext(ctx, `
		insert into collections (
			id,
			title,
			embedding_size,
			binary_filepath,
			created_at
		) values ($1, $2, $3, $4, $5);
	`,
		collection.ID,
		collection.Title,
		collection.EmbeddingSize,
		collection.BinaryFilepath,
		time.Now(),
	)

	if err != nil {
		return uuid.UUID{}, err
	}

	return collection.ID, nil
}

func (r *CollectionRepository) GetByTitle(ctx context.Context, title string) (*model.Collection, error) {
	row := r.conn.QueryRowContext(ctx, `
		select
			id,
			title,
			embedding_size,
			binary_filepath,
			created_at
		from collections
		where title = $1;
	`, title)

	var collection model.Collection
	err := row.Scan(
		&collection.ID,
		&collection.Title,
		&collection.EmbeddingSize,
		&collection.BinaryFilepath,
		&collection.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &collection, nil
}

func (r *CollectionRepository) CreateBinaryFile(ctx context.Context, title string) (string, error) {
	// TODO: форматировать путь хранения данных
	filepath := fmt.Sprintf("/Users/zrossiz/Desktop/vectorDB/mylvus/dbs/%s.dat", title)
	_, err := os.Create(filepath)
	if err != nil {
		return "", err
	}

	return filepath, nil
}
