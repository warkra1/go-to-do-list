package service

import (
	"to-do-list/app/infrastructure/password_hasher"
	"to-do-list/app/model"
	"to-do-list/app/repository"
)

type AuthService struct {
	currentUser model.User
	repository  repository.Repository[model.User]
}

func NewAuthService(repository repository.Repository[model.User]) *AuthService {
	return &AuthService{repository: repository}
}

func (s *AuthService) GetCurrentUser() model.User {
	return s.currentUser
}

func (s *AuthService) Login(login string, password string) *AuthError {
	user, err := s.repository.Read(login)
	if err != nil {
		return NewUserNotFoundAuthError()
	}

	if !password_hasher.CheckPassword(user, password) {
		return NewInvalidPasswordAuthError()
	}

	s.currentUser = user
	return nil
}

func (s *AuthService) Register(login string, password string) *AuthError {
	if s.exists(login) {
		return NewUserAlreadyExistsAuthError()
	}

	user := model.User{Login: login, Password: password_hasher.HashPassword(password)}
	s.repository.Write(login, user)
	s.currentUser = user
	return nil
}

func (s *AuthService) exists(login string) bool {
	_, err := s.repository.Read(login)
	return err == nil
}
