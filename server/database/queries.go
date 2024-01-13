package database

// Stored Procedures
const CREATE_TASK_SP = "CALL CreateTask(?, @taskId, @createdAt)"

// Simple queries
const GET_TASKS = "SELECT * FROM Task"
