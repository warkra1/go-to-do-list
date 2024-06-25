package repository

import (
	"encoding/json"
	"io"
	"os"
)

type Repository[T any] interface {
	Read(string) (T, *ModelNotFoundError)
	Write(string, T)
}

type FileRepository[T any] struct {
	path string
}

func NewFileRepository[T any](path string) *FileRepository[T] {
	return &FileRepository[T]{path: path}
}

func (f *FileRepository[T]) Read(id string) (T, *ModelNotFoundError) {
	var result T
	filepath := f.path + "/" + id + ".json"
	file, err := os.Open(filepath)
	if err != nil {
		return result, &ModelNotFoundError{}
	}
	defer file.Close()
	byteValue, _ := io.ReadAll(file)
	json.Unmarshal(byteValue, &result)
	return result, nil
}

func (f *FileRepository[T]) Write(id string, obj T) {
	filepath := f.path + "/" + id + ".json"
	file, _ := json.MarshalIndent(obj, "", "")
	os.WriteFile(filepath, file, 0644)
}
