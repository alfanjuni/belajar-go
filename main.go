package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"belajar-go/controllers"
	"belajar-go/models"
	"belajar-go/repositories"
	"belajar-go/services"
)

var DB *gorm.DB

// initDatabase initializes the database connection and migrates the Todo model
func initDatabase() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Construct the database connection string
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	// Open a connection to the database
	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if errDB != nil {
		log.Fatalf("Error connecting to database: %v", errDB)
	}

	// Automatically migrate the Todo model to keep the schema up to date
	err = DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Fatalf("Error migrating database: %v", err)
	}
}

func main() {
	initDatabase()

	// Initialize repository, service, and controller
	todoRepository := repositories.NewTodoRepository(DB)
	todoService := services.NewTodoService(todoRepository)
	todoController := controllers.NewTodoController(todoService)

	// Create a new router
	r := mux.NewRouter()

	// Define the API endpoints and their corresponding handler functions
	r.HandleFunc("/todos", todoController.GetTodos).Methods("GET")
	r.HandleFunc("/todos", todoController.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", todoController.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", todoController.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", todoController.DeleteTodo).Methods("DELETE")

	// Start the server on port 8080
	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
