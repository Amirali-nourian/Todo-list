package repository

import (
	"errors"
	"github.com/Amirali-nourian/Todo-list-golang/internal/domain"
	"gorm.io/gorm"
)

type TodoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *TodoRepository {
	return &TodoRepository{db: db}
}

func (r *TodoRepository) Create(todo *domain.Todo) error {
	return r.db.Create(todo).Error
}

func (r *TodoRepository) GetAll() ([]domain.Todo, error) {
	var todos []domain.Todo
	err := r.db.Find(&todos).Error
	return todos, err
}

func (r *TodoRepository) GetByID(id uint) (*domain.Todo, error) {
	var todo domain.Todo
	err := r.db.First(&todo, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("todo not found")
		}
		return nil, err
	}
	return &todo, nil
}

func (r *TodoRepository) Update(todo *domain.Todo) error {
	return r.db.Save(todo).Error
}

func (r *TodoRepository) Delete(id uint) error {
	result := r.db.Delete(&domain.Todo{}, id)
	if result.RowsAffected == 0 {
		return errors.New("todo not found")
	}
	return result.Error
}
