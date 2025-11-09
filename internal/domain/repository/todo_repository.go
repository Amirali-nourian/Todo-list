package repository

import "todo-list-golang/internal/domain/entity"

type TodoRepository interface {
    Save(todo *entity.Todo) error
    FindByID(id uint) (*entity.Todo, error)
    FindAll() ([]*entity.Todo, error)
    Delete(id uint) error
}