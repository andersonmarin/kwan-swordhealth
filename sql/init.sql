DROP TABLE IF EXISTS `tasks` CASCADE;

DROP TABLE IF EXISTS `users` CASCADE;

CREATE TABLE `users` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `username` VARCHAR(256) COLLATE utf8mb4_unicode_ci NOT NULL,
    `password` VARCHAR(256) COLLATE utf8mb4_unicode_ci NOT NULL,
    `role` ENUM('manager','technician') COLLATE utf8mb4_unicode_ci NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `users_username_IDX` (`username`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE `tasks` (
    `id` INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
    `user_id` INT(10) UNSIGNED NOT NULL,
    `summary` VARCHAR(2500) COLLATE utf8mb4_unicode_ci NOT NULL,
    `performed_at` TIMESTAMP NOT NULL,
    PRIMARY KEY (`id`),
    FOREIGN KEY (`user_id`) REFERENCES users(`id`),
    INDEX `tasks_performed_at_IDX` (`performed_at`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

INSERT INTO `users`(`id`, `username`, `password`, `role`)
VALUES (1, 'tech1', '$2a$12$aBDvyBvdKceGKcuyayLVd.9htwnGVLMK5r3zvZedDQKF3g28Xl6ey', 'technician'),
       (2, 'tech2', '$2a$12$AaOJi1d2Ey71Th2Ed4DJwOe5YuiW3N1HLyyepE2zeKEXUKeVYTrTa', 'technician'),
       (3, 'boss3', '$2a$12$2DlWesChA3PQ/JFBEpTaouq.rT.FJG.6grf1hxIpulslPW9rkoGXS', 'manager');