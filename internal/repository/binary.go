package repository

import (
	"encoding/binary"
	"io"
	"os"
)

type BinaryEmbeddingRepository struct {
	filepath string
	embLen   int
}

func NewBinaryEmbeddingRepository(filepath string, embLen int) *BinaryEmbeddingRepository {
	repo := BinaryEmbeddingRepository{
		filepath: filepath,
		embLen:   embLen,
	}

	return &repo
}

func (r *BinaryEmbeddingRepository) InsertAtOffset(embedding []float32, offset int64) (int64, error) {
	file, err := os.OpenFile(
		r.filepath,
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	ofst, err := file.Seek(offset, io.SeekStart)
	if err != nil {
		return 0, err
	}

	if err := binary.Write(file, binary.LittleEndian, embedding); err != nil {
		return 0, err
	}

	return ofst, nil
}

func (r *BinaryEmbeddingRepository) InsertAtEnd(embedding []float32) (int64, error) {
	file, err := os.OpenFile(
		r.filepath,
		os.O_RDWR|os.O_CREATE,
		0644,
	)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	offset, err := file.Seek(0, io.SeekEnd)
	if err != nil {
		return 0, err
	}

	if err := binary.Write(file, binary.LittleEndian, embedding); err != nil {
		return 0, err
	}

	return offset, nil
}

func (r *BinaryEmbeddingRepository) GetEmbeddingByOffset(offset int64) ([]float32, error) {
	file, err := os.Open(r.filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	embedding := make([]float32, r.embLen)

	_, err = file.Seek(offset, io.SeekStart)
	if err != nil {
		return nil, err
	}

	if err := binary.Read(file, binary.LittleEndian, embedding); err != nil {
		return nil, err
	}

	return embedding, nil
}
