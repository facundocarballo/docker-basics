CREATE DATABASE DOCKER_TODO_LIST;
USE DOCKER_TODO_LIST;

CREATE TABLE Task (
	id INT AUTO_INCREMENT PRIMARY KEY,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);