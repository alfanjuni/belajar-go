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
)

var DB *gorm.DB

func initDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))

	var errDB error
	DB, errDB = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if errDB != nil {
		log.Fatalf("Error connecting to database: %v", errDB)
	}

	DB.AutoMigrate(&models.Todo{})
}

func main() {
	initDatabase()
	controllers.Init(DB) // Initialize the controller with the DB instance

	r := mux.NewRouter()

	r.HandleFunc("/todos", controllers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", controllers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", controllers.GetTodo).Methods("GET")
	r.HandleFunc("/todos/{id}", controllers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", controllers.DeleteTodo).Methods("DELETE")

	log.Println("Server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
