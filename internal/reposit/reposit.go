package reposit

import (
	"github.com/jmoiron/sqlx"
	todo "github.com/zloyboy/todoapp"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Reposit struct {
	Authorization
	TodoList
	TodoItem
}

func NewReposit(db *sqlx.DB) *Reposit {
	return &Reposit{
		Authorization: NewAuthPostgres(db),
	}
}
