package main

import (
	"encoding/json"
	"go_server/internal/database"
	"go_server/internal/handlers"
	"go_server/internal/repositories"
	"go_server/internal/services"
	"go_server/internal/middleware"
	"github.com/joho/godotenv"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "go_server/docs"

	"log"
	"net/http"
)

func HealthHandler(w http.ResponseWriter, r *http.Request){
	response := map[string]string{
		"status": "OK",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main(){

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found")
	}

	port := os.Getenv("PORT")
	mongoURI := os.Getenv("MONGO_URI")
	dbName := os.Getenv("DB_NAME")
	database.ConnectMongo(mongoURI)

	taskRepo := repositories.NewMongoTaskRepository(dbName, "tasks")
	taskService := services.NewTaskService(taskRepo)
	taskHandler := handlers.NewTaskHandler(taskService)
	handlerWithMiddleware :=
		middleware.RecoveryMiddleware(
			middleware.LoggingMiddleware(taskHandler.TaskHandler),
		)
	http.HandleFunc("/health", HealthHandler)
	http.HandleFunc("/tasks", handlerWithMiddleware)
	http.HandleFunc("/tasks/", handlerWithMiddleware)
	http.Handle("/swagger/", httpSwagger.WrapHandler)


	log.Println("Server running on port: ", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}