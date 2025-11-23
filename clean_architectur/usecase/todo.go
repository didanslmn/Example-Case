package usecase

import (
	"clean_architectur/entity"
	"clean_architectur/repository"
)

type TodoUseCase struct {
	repo repository.TodoRepository
}

// constructor for TodoUseCase
func NewTodoUseCase(repo repository.TodoRepository) *TodoUseCase {
	return &TodoUseCase{repo: repo}
}

// method to create todo
func (u *TodoUseCase) CreateTodo(task string) error {
	return u.repo.Create(&entity.Todo{Task: task, Completed: false})
}

// method to get todo by id
func (u *TodoUseCase) GetTodoByID(id int) (*entity.Todo, error) {
	return u.repo.GetByID(id)
}

// method to get all todos
func (u *TodoUseCase) GetAllTodos() ([]*entity.Todo, error) {
	return u.repo.GetAll()
}
