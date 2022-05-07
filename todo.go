package todo

type TodoList struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Descript string `json:"descript"`
}

type TodoItem struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Descript string `json:"descript"`
	Done     bool   `json:"done"`
}

type UserList struct {
	Id     int
	UserId int
	ListId int
}

type ListItem struct {
	Id     int
	ListId int
	ItemId int
}
