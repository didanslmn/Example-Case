package repository

import "clean_architectur/entity"

type MemoryTodoRepository struct {
	todos []*entity.Todo
}

func NewMemoryTodoRepository() *MemoryTodoRepository {
	return &MemoryTodoRepository{todos: make([]*entity.Todo, 0)}
}

func (r *MemoryTodoRepository) Create(todo *entity.Todo) error {
	// Generate ID baru berdasarkan jumlah item yang ada + 1
	todo.ID = len(r.todos) + 1
	r.todos = append(r.todos, todo)
	return nil
}

func (r *MemoryTodoRepository) GetByID(id int) (*entity.Todo, error) {
	for _, todo := range r.todos {
		if todo.ID == id {
			return todo, nil
		}
	}
	return nil, nil
}
func (r *MemoryTodoRepository) GetAll() ([]*entity.Todo, error) {
	return r.todos, nil
}
