package controllers

import (
	"encoding/json"
	"net/http"

	"belajar-go/helpers"
	"belajar-go/models"
	"belajar-go/responses"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(db *gorm.DB) {
	DB = db
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	result := DB.Find(&todos)
	if result.Error != nil {
		helpers.RespondError(w, http.StatusInternalServerError, "Error fetching todos")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, todos, responses.Meta{Total: int(result.RowsAffected)}, "OK")
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = DB.Create(&todo).Error
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, "Error creating todo")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, todo, responses.Meta{Total: 1}, "Todo created")
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	err := DB.First(&todo, params["id"]).Error
	if err != nil {
		helpers.RespondError(w, http.StatusNotFound, "Todo not found")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, todo, responses.Meta{Total: 1}, "OK")
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	err := DB.First(&todo, params["id"]).Error
	if err != nil {
		helpers.RespondError(w, http.StatusNotFound, "Todo not found")
		return
	}

	err = json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		helpers.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = DB.Save(&todo).Error
	if err != nil {
		helpers.RespondError(w, http.StatusInternalServerError, "Error updating todo")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, todo, responses.Meta{Total: 1}, "Todo updated")
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	err := DB.Delete(&models.Todo{}, params["id"]).Error
	if err != nil {
		helpers.RespondError(w, http.StatusNotFound, "Todo not found")
		return
	}

	helpers.RespondJSON(w, http.StatusOK, nil, responses.Meta{Total: 0}, "Todo deleted")
}
