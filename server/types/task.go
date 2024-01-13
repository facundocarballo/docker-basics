package types

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	"github.com/facundocarballo/docker-basics/database"
)

type Task struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	CreatedAt   []uint8 `json:"created_at"`
}

func BodyToTrade(body []byte) *Task {
	if len(body) == 0 {
		return nil
	}

	var task Task
	err := json.Unmarshal(body, &task)
	if err != nil {
		return nil
	}

	return &task
}

func CreateTask(description string, db *sql.DB) *Task {
	statement, err := db.Prepare(database.CREATE_TASK_SP)
	if err != nil {
		fmt.Printf("Error preparing the Stored Procedure to create the new task.")
		return nil
	}
	defer statement.Close()

	_, err = statement.Exec(description)
	if err != nil {
		fmt.Printf("Error executing the Stored Procedure to create the new task.")
		return nil
	}

	var id int
	var createdAt []uint8
	err = db.QueryRow("SELECT @taskId, @createdAt").Scan(&id, &createdAt)
	if err != nil {
		fmt.Printf("Error selecting the tokenId and the createdAt OUT variables from the Stored Procedure.")
		return nil
	}

	return &Task{
		Id:          id,
		Description: description,
		CreatedAt:   createdAt,
	}
}

func GetTasks(db *sql.DB) []Task {
	rows, err := db.Query(database.GET_TASKS)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Description, &task.CreatedAt)
		if err != nil {
			log.Fatal(err)
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return tasks
}
