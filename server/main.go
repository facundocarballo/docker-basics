package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/facundocarballo/docker-basics/database"
	"github.com/facundocarballo/docker-basics/handlers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "¡Hello, World!")
}

func HandleTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	if r.Method == http.MethodGet {
		handlers.GetAllTasks(w, r, db)
		return
	}

	if r.Method == http.MethodPost {
		handlers.CreateTask(w, r, db)
		return
	}

	if r.Method == http.MethodDelete {
		handlers.DeleteTask(w, r, db)
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

	fmt.Println("Servidor escuchando en el puerto 8080...")
	http.ListenAndServe(":8080", nil)
}
