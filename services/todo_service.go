package services

import (
	"belajar-go/models"
	"belajar-go/repositories"
)

type TodoService interface {
	GetTodos() ([]models.Todo, error)
	GetTodoById(id uint) (models.Todo, error)
	CreateTodo(todo models.Todo) (models.Todo, error)
	UpdateTodo(todo models.Todo) (models.Todo, error)
	DeleteTodo(id uint) error
}

type todoService struct {
	repository repositories.TodoRepository
}

func NewTodoService(repo repositories.TodoRepository) TodoService {
	return &todoService{repo}
}

func (s *todoService) GetTodos() ([]models.Todo, error) {
	return s.repository.FindAll()
}

func (s *todoService) GetTodoById(id uint) (models.Todo, error) {
	return s.repository.FindById(id)
}

func (s *todoService) CreateTodo(todo models.Todo) (models.Todo, error) {
	return s.repository.Create(todo)
}

func (s *todoService) UpdateTodo(todo models.Todo) (models.Todo, error) {
	return s.repository.Update(todo)
}

func (s *todoService) DeleteTodo(id uint) error {
	return s.repository.Delete(id)
}
