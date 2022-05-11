package main

import (
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	todo "github.com/zloyboy/todoapp"
	"github.com/zloyboy/todoapp/config"
	"github.com/zloyboy/todoapp/internal/handler"
	"github.com/zloyboy/todoapp/internal/reposit"
	"github.com/zloyboy/todoapp/internal/service"
)

func main() {
	logrus.SetFormatter(&logrus.JSONFormatter{})
	config := config.ReadConfig("./config.json")

	db, err := reposit.NewPostgresDB(config)
	if err != nil {
		logrus.Fatal("init db error")
	}
	repos := reposit.NewReposit(db)
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(config.BindPort, handler.InitRoutes()); err != nil {
		logrus.Fatal("run error")
	}
}
