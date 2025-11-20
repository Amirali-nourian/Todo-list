package handler

import (
	"net/http"
	"strconv"
	"github.com/Amirali-nourian/Todo-list-golang/internal/domain"
	"github.com/Amirali-nourian/Todo-list-golang/internal/domain/service" // Ù…Ø·Ù…Ø¦Ù† Ø´Ùˆ Ú©Ù‡ Ø§ÛŒÙ† import Ø¯Ø±Ø³ØªÙ‡

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	Service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{Service: s}
}

// CreateTodo godoc
// @Summary Create a new todo
// @Description Create a new todo item
// @Tags todos
// @Accept  json
// @Produce  json
// @Param todo body domain.Todo true "Todo object"
// @Success 201 {object} domain.Todo
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var todo domain.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ø­Ø§Ù„Ø§ Ú©Ù„ Ø¢Ø¨Ø¬Ú©Øª todo Ø±Ùˆ Ø¨Ù‡ Ø³Ø±ÙˆÛŒØ³ Ù…ÛŒâ€ŒÙØ±Ø³ØªÛŒÙ…
	if err := h.Service.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create todo"})
		return
	}

	c.JSON(http.StatusCreated, todo)
}

// GetAllTodos godoc
// @Summary Get all todos
// @Description Get a list of all todo items
// @Tags todos
// @Produce  json
// @Success 200 {array} domain.Todo
// @Failure 500 {object} map[string]string
// @Router /todos [get]
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	// Ø§Ø³Ù… ØªØ§Ø¨Ø¹ Ø¯Ø± Ø³Ø±ÙˆÛŒØ³ GetAllTodos Ø¨ÙˆØ¯
	todos, err := h.Service.GetAllTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get todos"})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// GetTodoByID godoc
// @Summary Get a todo by ID
// @Description Get a single todo item by its ID
// @Tags todos
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 200 {object} domain.Todo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [get]
func (h *TodoHandler) GetTodoByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Ø§Ø³Ù… ØªØ§Ø¨Ø¹ Ø¯Ø± Ø³Ø±ÙˆÛŒØ³ GetTodoByID Ø¨ÙˆØ¯
	todo, err := h.Service.GetTodoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// UpdateTodo godoc
// @Summary Update a todo
// @Description Update an existing todo item by ID
// @Tags todos
// @Accept  json
// @Produce  json
// @Param id path int true "Todo ID"
// @Param todo body domain.Todo true "Todo object"
// @Success 200 {object} domain.Todo
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /todos/{id} [put]
func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	// Ø§ÙˆÙ„ Ú†Ú© Ù…ÛŒâ€ŒÚ©Ù†ÛŒÙ… todo ÙˆØ¬ÙˆØ¯ Ø¯Ø§Ø±Ù‡
	todo, err := h.Service.GetTodoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// Ø¯ÛŒØªØ§ÛŒ Ø¬Ø¯ÛŒØ¯ Ø±Ùˆ Ø§Ø² body Ù…ÛŒâ€ŒØ®ÙˆÙ†ÛŒÙ… Ùˆ Ø±ÙˆÛŒ Ø¢Ø¨Ø¬Ú©Øª Ù‚Ø¨Ù„ÛŒ Ù…ÛŒâ€ŒØ±ÛŒØ²ÛŒÙ…
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Ù…Ø·Ù…Ø¦Ù† Ù…ÛŒâ€ŒØ´ÛŒÙ… ID Ø¹ÙˆØ¶ Ù†Ø´Ø¯Ù‡ Ø¨Ø§Ø´Ù‡
	todo.ID = uint(id)

	// Ø³Ø±ÙˆÛŒØ³ Ù…Ø§ ÛŒÚ© Ø¢Ø¨Ø¬Ú©Øª Ú©Ø§Ù…Ù„ todo Ù…ÛŒâ€ŒØ®ÙˆØ§Ø³Øª
	if err := h.Service.UpdateTodo(todo); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

// DeleteTodo godoc
// @Summary Delete a todo
// @Description Delete a todo item by ID
// @Tags todos
// @Produce  json
// @Param id path int true "Todo ID"
// @Success 204 {object} nil
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /todos/{id} [delete]
func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	if err := h.Service.DeleteTodo(uint(id)); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found or failed to delete"})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
