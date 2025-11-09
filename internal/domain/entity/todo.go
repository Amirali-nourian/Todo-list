package entity

import "time"

type Todo struct {
    ID          uint      `json:"id"`
    Title       string    `json:"title"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
}

func NewTodo(title string) *Todo {
    return &Todo{
        Title:     title,
        Completed: false,
        CreatedAt: time.Now(),
    }
}