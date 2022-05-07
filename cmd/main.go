package main

import (
	"log"

	todo "github.com/zloyboy/todoapp"
	"github.com/zloyboy/todoapp/config"
	"github.com/zloyboy/todoapp/internal/handler"
	"github.com/zloyboy/todoapp/internal/reposit"
	"github.com/zloyboy/todoapp/internal/service"
)

func main() {
	config := config.ReadConfig("./config.json")
	repos := reposit.NewReposit()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(config.BindPort, handler.InitRoutes()); err != nil {
		log.Fatal("run error")
	}
}
