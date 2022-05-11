package reposit

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"github.com/zloyboy/todoapp/config"
)

const (
	userTable     = "users"
	todoListTable = "todo_list"
	userListTable = "user_list"
	todoItemTable = "todo_item"
	listItemTable = "list_item"
)

func NewPostgresDB(cfg config.Config) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLmode))
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	logrus.Print("db ping ok")
	return db, nil
}
