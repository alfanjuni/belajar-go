package controllers

import (
	"encoding/json"
	"net/http"

	"belajar-go/models"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init(db *gorm.DB) {
	DB = db
}

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	DB.Create(&todo)
	json.NewEncoder(w).Encode(todo)
}

func GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	if err := DB.First(&todo, params["id"]).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var todo models.Todo
	if err := DB.First(&todo, params["id"]).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	DB.Save(&todo)
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if err := DB.Delete(&models.Todo{}, params["id"]).Error; err != nil {
		http.Error(w, "Todo not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"message": "Todo deleted"})
}
