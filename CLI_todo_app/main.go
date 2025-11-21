package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Todo struct {
	ID        int       `json:"id`
	Task      string    `json:"task"`
	Completed bool      `json:"completed`
	CreatedAt time.Time `json:"created_at"`
}

// save todo ke file json

func saveTodo(todos []Todo, filename string) error {
	data, err := json.MarshalIndent(todos, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

// load todo dari file

func loadTodo(filename string) ([]Todo, error) {
	var todos []Todo
	data, err := os.ReadFile(filename)
	if err != nil {
		if os.IsNotExist(err) {
			return []Todo{}, nil // file belum ada return emty slice
		}
		return nil, err
	}
	json.Unmarshal(data, &todos)
	return todos, nil
}

func main() {
	filename := "todos.json"

	//load existing todos
	todos, _ := loadTodo(filename)

	// add new todos
	newTodo := Todo{
		ID:        len(todos) + 1,
		Task:      "Belajar Golang",
		Completed: true,
		CreatedAt: time.Now(),
	}
	todos = append(todos, newTodo)

	// save file
	if err := saveTodo(todos, filename); err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Todo berhasil ditambahkan")

	// print all todos
	for _, todo := range todos {
		status := "belum selesai"
		if todo.Completed {
			status = "selesai"
		}
		fmt.Printf("%d. %s - %s (dibuat:):%s\n", todo.ID, todo.Task, status, todo.CreatedAt.Format("02 Jan 15:04"))
	}
}
