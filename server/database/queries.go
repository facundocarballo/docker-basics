package database

// Stored Procedures
const CREATE_TASK_SP = "CALL CreateTask(?, @taskId, @createdAt)"
const DELETE_TASK_SP = "CALL DeleteTask(?, @exist)"

// Simple queries
const Q_GET_TASKS = "SELECT * FROM Task"
