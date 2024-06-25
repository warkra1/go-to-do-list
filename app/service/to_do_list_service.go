package service

import (
	"to-do-list/app/model"
	"to-do-list/app/repository"
)

type ToDoListService struct {
	repository *repository.ToDoListRepository
}

func NewToDoListService(repository *repository.ToDoListRepository) *ToDoListService {
	return &ToDoListService{repository: repository}
}

func (s *ToDoListService) createList(user model.User) model.ToDoList {
	list := model.ToDoList{Items: []model.ToDoItem{}}
	s.repository.Write(user.Login, list)
	return list
}

func (s *ToDoListService) GetList(user model.User) model.ToDoList {
	list, err := s.repository.Read(user.Login)
	if err == nil {
		return list
	}
	return s.createList(user)
}

func (s *ToDoListService) CreateItem(user model.User, item model.ToDoItem, number int) model.ToDoList {
	list := s.GetList(user)
	if number > len(list.Items) {
		number = len(list.Items)
	}

	if number == len(list.Items) {
		list.Items = append(list.Items, item)
	} else {
		list.Items = append(list.Items[:number], append([]model.ToDoItem{item}, list.Items[number:]...)...)
	}
	s.repository.Write(user.Login, list)
	return list
}

func (s *ToDoListService) UpdateItem(user model.User, item model.ToDoItem, number int) model.ToDoList {
	list := s.GetList(user)
	if number < len(list.Items) {
		list.Items[number] = item
	}
	s.repository.Write(user.Login, list)
	return list
}

func (s *ToDoListService) DeleteItem(user model.User, number int) model.ToDoList {
	list := s.GetList(user)
	if number < len(list.Items) {
		copy(list.Items[number:], list.Items[number+1:])
		list.Items = list.Items[:len(list.Items)-1]
	}
	s.repository.Write(user.Login, list)
	return list
}
