package repository

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type BinaryEmbeddingRepository struct{}

func NewBinaryEmbeddingRepository() *BinaryEmbeddingRepository {
	return &BinaryEmbeddingRepository{}
}

func (r *BinaryEmbeddingRepository) InsertAtOffset(filepath string, embedding []float32, offset int64) error {
	file, err := os.OpenFile(
		filepath,
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return err
	}

	if err := binary.Write(file, binary.LittleEndian, embedding); err != nil {
		return err
	}

	return nil
}

func (r *BinaryEmbeddingRepository) InsertAtEnd(filepath string, embedding []float32) (int64, error) {
	file, err := os.OpenFile(
		filepath,
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	_, err = file.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}

	if err := binary.Write(file, binary.LittleEndian, embedding); err != nil {
		return 0, err
	}

	endOffset, err := file.Seek(0, io.SeekCurrent)
	if err != nil {
		return 0, err
	}

	return endOffset, nil
}

func (r *BinaryEmbeddingRepository) GetEmbeddingByOffset(filepath string, offset, embLen int64) ([]float32, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	embedding := make([]float32, embLen)

	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	if err := binary.Read(file, binary.LittleEndian, embedding); err != nil {
		return nil, err
	}

	return embedding, nil
}

func (r *BinaryEmbeddingRepository) Truncate(filepath string, size int64) error {
	if err := os.Truncate(filepath, size); err != nil {
		return fmt.Errorf("truncate vector file to %d bytes: %w", size, err)
	}

	return nil
}
