package types

import (
	"database/sql"
	"encoding/json"
	"fmt"

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
		fmt.Printf("Error selecting the taskId and the createdAt OUT variables from the Stored Procedure.")
		return nil
	}

	return &Task{
		Id:          id,
		Description: description,
		CreatedAt:   createdAt,
	}
}

func GetTasks(db *sql.DB) []Task {
	rows, err := db.Query(database.Q_GET_TASKS)
	if err != nil {
		fmt.Printf("Error quering tasks: %s\n", err.Error())
		return nil
	}
	defer rows.Close()

	var tasks []Task
	for rows.Next() {
		var task Task
		err := rows.Scan(&task.Id, &task.Description, &task.CreatedAt)
		if err != nil {
			fmt.Printf("Error scanning rows: %s\n", err.Error())
			return nil
		}
		tasks = append(tasks, task)
	}

	if err := rows.Err(); err != nil {
		fmt.Printf("Error rows: %s\n", err.Error())
		return nil
	}

	return tasks
}

func DeleteTask(id int, db *sql.DB) bool {
	statement, err := db.Prepare(database.DELETE_TASK_SP)
	if err != nil {
		fmt.Printf("Error preparing the Stored Procedure to delete a task.")
		return false
	}
	defer statement.Close()

	_, err = statement.Exec(id)
	if err != nil {
		fmt.Printf("Error executing the Stored Procedure to delete the task.")
		return false
	}

	var exist int
	err = db.QueryRow("SELECT @exist").Scan(&exist)
	if err != nil {
		fmt.Printf("Error selecting the exist OUT variable from the Stored Procedure.")
		return false
	}

	return exist > 0
}
