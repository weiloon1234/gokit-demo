CREATE TABLE `users` (
    `id` bigint UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `email_verified_at` timestamp NULL DEFAULT NULL,
    `password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `password2` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
    `country_id` bigint UNSIGNED DEFAULT NULL,
    `contact_country_id` bigint UNSIGNED DEFAULT NULL,
    `contact_number` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `full_contact_number` varchar(24) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `introducer_user_id` bigint UNSIGNED DEFAULT NULL,
    `lang` varchar(12) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'en',
    `avatar` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci,
    `credit_1` decimal(18,5) NOT NULL DEFAULT '0.00000',
    `credit_2` decimal(18,5) NOT NULL DEFAULT '0.00000',
    `credit_3` decimal(18,5) NOT NULL DEFAULT '0.00000',
    `credit_4` decimal(18,5) NOT NULL DEFAULT '0.00000',
    `credit_5` decimal(18,5) NOT NULL DEFAULT '0.00000',
    `bank_id` bigint UNSIGNED DEFAULT NULL,
    `bank_account_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `bank_account_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `national_id` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci DEFAULT NULL,
    `first_login` tinyint UNSIGNED NOT NULL DEFAULT '0',
    `ban_until` timestamp NULL DEFAULT NULL,
    `new_login_at` datetime DEFAULT NULL,
    `last_login_at` datetime DEFAULT NULL,
    `unilevel` bigint UNSIGNED DEFAULT NULL,
    `created_at` timestamp NULL DEFAULT NULL,
    `updated_at` timestamp NULL DEFAULT NULL,
    `deleted_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE `users`
    -- Add indexes for faster lookups
    ADD INDEX `index_username` (`username`),
    ADD INDEX `index_email` (`email`),
    ADD INDEX `fk_country_id` (`country_id`),
    ADD INDEX `fk_contact_country_id` (`contact_country_id`),
    ADD INDEX `fk_introducer_user_id` (`introducer_user_id`),
    ADD INDEX `fk_bank_id` (`bank_id`),

    -- Add foreign key constraints
    ADD CONSTRAINT `fk_users_country_id` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_users_contact_country_id` FOREIGN KEY (`contact_country_id`) REFERENCES `countries` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_users_introducer_user_id` FOREIGN KEY (`introducer_user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_users_bank_id` FOREIGN KEY (`bank_id`) REFERENCES `banks` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
;