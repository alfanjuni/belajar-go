CREATE TABLE IF NOT EXISTS `todos` (
    `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
    `deleted_at` datetime(3) DEFAULT NULL,
    `title` varchar(255) DEFAULT NULL,
    `description` text,
    `completed` tinyint(1) DEFAULT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_todos_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


INSERT INTO todos (created_at, updated_at, title, description, completed) 
VALUES (NOW(), NOW(), 'Learn Go', 'Learn how to create a REST API using Go', false);


SELECT * from todos;
