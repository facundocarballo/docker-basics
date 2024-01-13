package handlers

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"

	"github.com/facundocarballo/docker-basics/types"
)

func GetAllTasks(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	tasks := types.GetTasks(db)
	if tasks == nil {
		http.Error(w, "Error getting tasks.", http.StatusNotFound)
		return
	}

	jsonData, err := json.Marshal(tasks)
	if err != nil {
		http.Error(w, "Error JSON converting the tasks to JSON.", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func CreateTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading the body of request.", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	task := types.BodyToTrade(body)
	task = types.CreateTask(task.Description, db)

	if task == nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Error creating the task in the database. " + err.Error()))
		return
	}

	jsonData, err := json.Marshal(task)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Error converting the task to JSON. " + err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func DeleteTask(w http.ResponseWriter, r *http.Request, db *sql.DB) {

}
