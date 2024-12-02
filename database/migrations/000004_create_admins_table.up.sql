CREATE TABLE `admins` (
    `id` bigint UNSIGNED NOT NULL PRIMARY KEY,
    `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `email_verified_at` timestamp NULL DEFAULT NULL,
    `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `password2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `lang` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'en',
    `avatar` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `type` tinyint UNSIGNED NOT NULL DEFAULT '0',
    `created_at` timestamp NULL DEFAULT NULL COMMENT 'Record creation timestamp',
    `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Record deleted timestamp'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE `admins`
    ADD INDEX `index_username` (`username`),
    ADD INDEX `index_email` (`email`)
;