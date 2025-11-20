package service

import "todo-list-golang/internal/domain"

// این interface به Service می‌گه که از Repository چی می‌خواد
// این باید با فایل repository.go ما مچ باشه
type TodoRepository interface {
	Create(todo *domain.Todo) error
	GetAll() ([]domain.Todo, error)
	GetByID(id uint) (*domain.Todo, error)
	Update(todo *domain.Todo) error
	Delete(id uint) error
}

type TodoService struct {
	repo TodoRepository
}

func NewTodoService(repo TodoRepository) *TodoService {
	return &TodoService{repo: repo}
}

func (s *TodoService) CreateTodo(todo *domain.Todo) error {
	return s.repo.Create(todo)
}

func (s *TodoService) GetAllTodos() ([]domain.Todo, error) {
	return s.repo.GetAll()
}

func (s *TodoService) GetTodoByID(id uint) (*domain.Todo, error) {
	return s.repo.GetByID(id)
}

func (s *TodoService) UpdateTodo(todo *domain.Todo) error {
	return s.repo.Update(todo)
}

func (s *TodoService) DeleteTodo(id uint) error {
	return s.repo.Delete(id)
}