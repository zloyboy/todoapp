package reposit

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

func NewReposit() *Reposit {
	return &Reposit{}
}
