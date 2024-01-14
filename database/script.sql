CREATE TABLE `Task` (
	id INT AUTO_INCREMENT PRIMARY KEY,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

DELIMITER //
	CREATE PROCEDURE `CreateTask`(IN description VARCHAR(255), OUT taskId INT,  OUT createdTimestamp TIMESTAMP)
	BEGIN
		INSERT INTO Task (description) VALUES (description);
		SELECT LAST_INSERT_ID() INTO taskId;
		SELECT created_at INTO createdTimestamp FROM Task WHERE id = taskId;
	END //
DELIMITER ;

DELIMITER //
	CREATE PROCEDURE `DeleteTask`(IN taskId INT, OUT exist INT)
	BEGIN
        SELECT COUNT(*) INTO exist FROM Task WHERE id = taskId;
        IF exist > 0 THEN
			DELETE FROM Task WHERE id = taskId;
		END IF;
	END //
DELIMITER ;
