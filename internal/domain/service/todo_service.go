package service

import (
    "todo-list-golang/internal/domain/entity"
    "todo-list-golang/internal/domain/repository"
)

type TodoService struct {
    Repo repository.TodoRepository
}

func NewTodoService(repo repository.TodoRepository) *TodoService {
    return &TodoService{Repo: repo}
}

func (s *TodoService) CreateTodo(title string) (*entity.Todo, error) {
    todo := entity.NewTodo(title)
    err := s.Repo.Save(todo)
    if err != nil {
        return nil, err
    }
    return todo, nil
}

func (s *TodoService) GetTodo(id uint) (*entity.Todo, error) {
    return s.Repo.FindByID(id)
}

func (s *TodoService) UpdateTodo(id uint, completed bool) error {
    todo, err := s.Repo.FindByID(id)
    if err != nil {
        return err
    }
    todo.Completed = completed
    return s.Repo.Save(todo)
}

func (s *TodoService) DeleteTodo(id uint) error {
    todo, err := s.Repo.FindByID(id)
    if err != nil {
        return err
    }
    return s.Repo.Delete(todo.ID)
}

func (s *TodoService) ListTodos() ([]*entity.Todo, error) {
    return s.Repo.FindAll()
}
