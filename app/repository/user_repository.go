package repository

import (
	"to-do-list/app/model"
)

type UserRepository struct {
	FileRepository[model.User]
}

func (r *UserRepository) Exists(login string) bool {
	_, err := r.Read(login)
	return err == nil
}

func NewUserRepository(path string) *UserRepository {
	return &UserRepository{FileRepository: FileRepository[model.User]{path: path}}
}
