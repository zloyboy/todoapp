package main

import (
	"log"

	todo "github.com/zloyboy/todoapp"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatal("run error")
	}
}
