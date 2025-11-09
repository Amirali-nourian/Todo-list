package handler

import (
	"net/http"
	"strconv"
	"todo-list-golang/internal/domain/service"

	"github.com/gin-gonic/gin"
)

// Request structs for Swagger
type createReq struct {
	Title string `json:"title" binding:"required" example:"Buy milk"`
}
type updateReq struct {
	Completed bool `json:"completed" example:"true"`
}

// === STRUCT & CONSTRUCTOR ===
type TodoHandler struct {
	Service *service.TodoService
}

func NewTodoHandler(s *service.TodoService) *TodoHandler {
	return &TodoHandler{Service: s}
}

// === SWAGGER ENDPOINTS ===

// CreateTodo godoc
// @Summary      Create a new todo
// @Description  Create a todo item with a title
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        body  body      handler.createReq  true  "Todo title"
// @Success      201   {object}  entity.Todo
// @Failure      400   {object}  map[string]string  "Invalid input"
// @Failure      500   {object}  map[string]string  "Server error"
// @Router       /api/v1/todos [post]
func (h *TodoHandler) Create(c *gin.Context) {
	var req struct {
		Title string `json:"title" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo, err := h.Service.CreateTodo(req.Title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

// GetAllTodos godoc
// @Summary      Get all todos
// @Description  Retrieve list of all todo items
// @Tags         todos
// @Produce      json
// @Success      200  {array}   entity.Todo
// @Failure      500  {object}  map[string]string  "Server error"
// @Router       /api/v1/todos [get]
func (h *TodoHandler) GetAll(c *gin.Context) {
	todos, err := h.Service.ListTodos()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todos)
}

// GetOneTodo godoc
// @Summary      Get a single todo
// @Description  Get todo by ID
// @Tags         todos
// @Produce      json
// @Param        id   path      uint  true  "Todo ID"
// @Success      200  {object}  entity.Todo
// @Failure      404  {object}  map[string]string  "Todo not found"
// @Router       /api/v1/todos/{id} [get]
func (h *TodoHandler) GetOne(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)
	todo, err := h.Service.GetTodo(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, todo)
}

// UpdateTodo godoc
// @Summary      Update a todo
// @Description  Mark a todo as completed or not
// @Tags         todos
// @Accept       json
// @Produce      json
// @Param        id     path      uint  true  "Todo ID"
// @Param        body   body      handler.updateReq  true  "Completed status"
// @Success      200    {object}  map[string]string  "Updated"
// @Failure      400    {object}  map[string]string  "Invalid input"
// @Failure      404    {object}  map[string]string  "Todo not found"
// @Router       /api/v1/todos/{id} [put]
func (h *TodoHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	var req struct {
		Completed bool `json:"completed"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := h.Service.UpdateTodo(uint(id), req.Completed)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

// DeleteTodo godoc
// @Summary      Delete a todo
// @Description  Remove a todo by ID
// @Tags         todos
// @Produce      json
// @Param        id   path      uint  true  "Todo ID"
// @Success      200  {object}  map[string]string  "Deleted"
// @Failure      404  {object}  map[string]string  "Todo not found"
// @Router       /api/v1/todos/{id} [delete]
func (h *TodoHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, _ := strconv.ParseUint(idStr, 10, 32)

	err := h.Service.DeleteTodo(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}