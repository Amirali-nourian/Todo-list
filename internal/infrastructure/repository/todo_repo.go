package repository

import (
    "errors"
    "sync"
    "todo-list-golang/internal/domain/entity"
)

var ErrNotFound = errors.New("todo not found")

type InMemoryTodoRepo struct {
    todos map[uint]*entity.Todo
    mu    sync.RWMutex
    nextID uint
}

func NewInMemoryTodoRepo() *InMemoryTodoRepo {
    return &InMemoryTodoRepo{
        todos:  make(map[uint]*entity.Todo),
        nextID: 1,
    }
}

func (r *InMemoryTodoRepo) Save(todo *entity.Todo) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    if todo.ID == 0 {
        todo.ID = r.nextID
        r.nextID++
    }
    r.todos[todo.ID] = todo
    return nil
}

func (r *InMemoryTodoRepo) FindByID(id uint) (*entity.Todo, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    todo, exists := r.todos[id]
    if !exists {
        return nil, ErrNotFound
    }
    return todo, nil
}

func (r *InMemoryTodoRepo) FindAll() ([]*entity.Todo, error) {
    r.mu.RLock()
    defer r.mu.RUnlock()
    todos := make([]*entity.Todo, 0, len(r.todos))
    for _, todo := range r.todos {
        todos = append(todos, todo)
    }
    return todos, nil
}

func (r *InMemoryTodoRepo) Delete(id uint) error {
    r.mu.Lock()
    defer r.mu.Unlock()
    if _, exists := r.todos[id]; !exists {
        return ErrNotFound
    }
    delete(r.todos, id)
    return nil
}