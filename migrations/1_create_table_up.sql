CREATE TABLE `user` (
                        `user_id` INT AUTO_INCREMENT PRIMARY KEY,
                        `email` VARCHAR(255) UNIQUE NOT NULL,
                        `password` VARCHAR(255) NOT NULL,
                        `role` VARCHAR(50) NOT NULL,
                        `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                        `updated_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
