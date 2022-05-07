package service

import "github.com/zloyboy/todoapp/internal/reposit"

type Authorization interface {
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
	return &Service{}
}
