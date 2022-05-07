package main

import (
	"log"

	todo "github.com/zloyboy/todoapp"
	"github.com/zloyboy/todoapp/internal/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8000", handler.InitRoutes()); err != nil {
		log.Fatal("run error")
	}
}
