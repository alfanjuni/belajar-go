package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"belajar-go/helpers"
	"belajar-go/models"
	"belajar-go/services"
	"belajar-go/responses"

	"github.com/gorilla/mux"
)

type TodoController struct {
	service services.TodoService
}

func NewTodoController(service services.TodoService) *TodoController {
	return &TodoController{service}
}

func (c *TodoController) GetTodos(w http.ResponseWriter, r *http.Request) {
	todos, err := c.service.GetTodos()
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, "Error fetching todos")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, todos, responses.Meta{Total: len(todos)}, "OK")
}

func (c *TodoController) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	todo, err = c.service.CreateTodo(todo)
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, "Error creating todo")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, todo, responses.Meta{Total: 1}, "Todo created")
}

func (c *TodoController) GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	todo, err := c.service.GetTodoById(uint(id))
	if err != nil {
		helpers.RespondError(w, http.StatusNotFound, "Todo not found")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, todo, responses.Meta{Total: 1}, "OK")
}

func (c *TodoController) UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	todo.ID = uint(id)
	todo, err = c.service.UpdateTodo(todo)
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, "Error updating todo")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, todo, responses.Meta{Total: 1}, "Todo updated")
}

func (c *TodoController) DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	err := c.service.DeleteTodo(uint(id))
	if err != nil {
		helpers.RespondError(w, http.StatusNotFound, "Todo not found")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, nil, responses.Meta{Total: 0}, "Todo deleted")
}
