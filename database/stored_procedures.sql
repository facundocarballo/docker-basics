CREATE DATABASE DOCKER_TODO_LIST;
USE DOCKER_TODO_LIST;

DROP PROCEDURE CreateTask;

DELIMITER //
	CREATE PROCEDURE CreateTask(IN description VARCHAR(255), OUT taskId INT,  OUT createdTimestamp TIMESTAMP)
	BEGIN
		INSERT INTO Task (description) VALUES (description);
		SELECT LAST_INSERT_ID() INTO taskId;
		SELECT created_at INTO createdTimestamp FROM Task WHERE id = taskId;
	END //
DELIMITER ;