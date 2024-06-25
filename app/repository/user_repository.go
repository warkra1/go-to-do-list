package repository

import (
	"to-do-list/app/model"
)

type UserRepository struct {
	FileRepository[model.User]
}

func NewUserRepository(path string) *UserRepository {
	return &UserRepository{FileRepository: FileRepository[model.User]{path: path}}
}
