package main

import (
	"fmt"
	"log"

	"clean_architectur/repository"
	uscase "clean_architectur/usecase"
)

func main() {
	// dependsi injection
	repo := repository.NewMemoryTodoRepository()
	usecase := uscase.NewTodoUseCase(repo)

	// gunakan usecase
	if err := usecase.CreateTodo("Belajar Clean Architecture"); err != nil {
		log.Fatalf("Gagal membuat todo: %v", err)
	}
	if err := usecase.CreateTodo("Mengerjakan tugas golang"); err != nil {
		log.Fatalf("Gagal membuat todo: %v", err)
	}

	// Contoh mengambil semua todos untuk verifikasi
	todos, err := usecase.GetAllTodos()
	if err != nil {
		log.Fatalf("Gagal mengambil todos: %v", err)
	}
	fmt.Println("Semua Todos:", todos)
	for _, todo := range todos {
		fmt.Printf("ID: %d, Task: %s, Completed: %v\n", todo.ID, todo.Task, todo.Completed)
	}

	// Contoh mengambil todo berdasarkan ID
	todo, err := usecase.GetTodoByID(1)
	if err != nil {
		log.Fatalf("Gagal mengambil todo: %v", err)
	}
	fmt.Printf("Todo dengan ID 1: %+v\n", todo)
}
