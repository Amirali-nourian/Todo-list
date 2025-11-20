package service

import "github.com/Amirali-nourian/Todo-list-golang/internal/domain"

// Ø§ÛŒÙ† interface Ø¨Ù‡ Service Ù…ÛŒâ€ŒÚ¯Ù‡ Ú©Ù‡ Ø§Ø² Repository Ú†ÛŒ Ù…ÛŒâ€ŒØ®ÙˆØ§Ø¯
// Ø§ÛŒÙ† Ø¨Ø§ÛŒØ¯ Ø¨Ø§ ÙØ§ÛŒÙ„ repository.go Ù…Ø§ Ù…Ú† Ø¨Ø§Ø´Ù‡
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
