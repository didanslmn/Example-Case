package repository

import "clean_architectur/entity"

type TodoRepository interface {
	Create(todo *entity.Todo) error
	GetByID(id int) (*entity.Todo, error)
	GetAll() ([]*entity.Todo, error)
}
