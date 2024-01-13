package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/facundocarballo/docker-basics/database"
	"github.com/facundocarballo/docker-basics/handler"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/handlers"
	"github.com/joho/godotenv"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hello, World!")
}

func HandleTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodGet {
		handler.GetAllTasks(w, r, db)
		return
	}

	if r.Method == http.MethodPost {
		handler.CreateTask(w, r, db)
		return
	}

	if r.Method == http.MethodDelete {
		handler.DeleteTask(w, r, db)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
}

func main() {
	// Load enviroment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading the .env file")
	}

	// Get DB Instance
	db, err := sql.Open("mysql", database.GetDSN())
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		HelloWorld(w, r)
	})

	http.HandleFunc("/task", func(w http.ResponseWriter, r *http.Request) {
		HandleTasks(w, r, db)
	})

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "OPTIONS", "PUT", "PATCH", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
		handlers.AllowCredentials(),
	)

	fmt.Println("Servidor escuchando en el puerto 8080...")
	http.ListenAndServe(":8080", corsMiddleware(http.DefaultServeMux))
}
