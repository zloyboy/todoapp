package main

import (
	"log"

	todo "github.com/zloyboy/todoapp"
	"github.com/zloyboy/todoapp/internal/handler"
	"github.com/zloyboy/todoapp/internal/reposit"
	"github.com/zloyboy/todoapp/internal/service"
)

func main() {
	repos := reposit.NewReposit()
	services := service.NewService(repos)
	handler := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatal("run error")
	}
}
