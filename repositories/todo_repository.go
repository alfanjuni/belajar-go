package repositories

import (
	"belajar-go/models"

	"gorm.io/gorm"
)

type TodoRepository interface {
	FindAll() ([]models.Todo, error)
	FindById(id uint) (models.Todo, error)
	Create(todo models.Todo) (models.Todo, error)
	Update(todo models.Todo) (models.Todo, error)
	Delete(id uint) error
}

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) TodoRepository {
	return &todoRepository{db}
}

func (r *todoRepository) FindAll() ([]models.Todo, error) {
	var todos []models.Todo
	result := r.db.Find(&todos)
	return todos, result.Error
}

func (r *todoRepository) FindById(id uint) (models.Todo, error) {
	var todo models.Todo
	result := r.db.First(&todo, id)
	return todo, result.Error
}

func (r *todoRepository) Create(todo models.Todo) (models.Todo, error) {
	result := r.db.Create(&todo)
	return todo, result.Error
}

func (r *todoRepository) Update(todo models.Todo) (models.Todo, error) {
	var existingTodo models.Todo
	if err := r.db.First(&existingTodo, todo.ID).Error; err != nil {
		return todo, err
	}

	// Copy CreatedAt to prevent it from being overwritten
	todo.CreatedAt = existingTodo.CreatedAt

	result := r.db.Save(&todo)
	return todo, result.Error
}

func (r *todoRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Todo{}, id)
	return result.Error
}
