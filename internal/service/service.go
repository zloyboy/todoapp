package service

import (
	todo "github.com/zloyboy/todoapp"
	"github.com/zloyboy/todoapp/internal/reposit"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *reposit.Reposit) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
