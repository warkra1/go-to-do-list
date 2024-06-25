package repository

import "to-do-list/app/model"

type ToDoListRepository struct {
	FileRepository[model.ToDoList]
}

func NewToDoListRepository(path string) *ToDoListRepository {
	return &ToDoListRepository{FileRepository: FileRepository[model.ToDoList]{path: path}}
}
