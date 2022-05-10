package reposit

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Reposit{}
}
