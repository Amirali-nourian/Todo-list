package handler

import (
	"net/http"
	"strconv"
	"todo-list-golang/internal/domain"
	"todo-list-golang/internal/domain/service" // مطمئن شو که این import درسته

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

	// حالا کل آبجکت todo رو به سرویس می‌فرستیم
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
	// اسم تابع در سرویس GetAllTodos بود
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

	// اسم تابع در سرویس GetTodoByID بود
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

	// اول چک می‌کنیم todo وجود داره
	todo, err := h.Service.GetTodoByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	// دیتای جدید رو از body می‌خونیم و روی آبجکت قبلی می‌ریزیم
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// مطمئن می‌شیم ID عوض نشده باشه
	todo.ID = uint(id)

	// سرویس ما یک آبجکت کامل todo می‌خواست
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