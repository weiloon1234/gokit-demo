CREATE TABLE `banks` (
    `id` bigint UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `name_en` VARCHAR(255) NOT NULL COMMENT 'Bank name in english',
    `name_zh` VARCHAR(255) NOT NULL COMMENT 'Bank name in chinese',
    `country_id` bigint UNSIGNED DEFAULT NULL COMMENT 'Country of the bank',
    `created_at` timestamp NULL DEFAULT NULL COMMENT 'Record creation timestamp',
    `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Record deleted timestamp'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE `banks`
    ADD INDEX `index_country_id` (`country_id`),

    ADD CONSTRAINT `fk_banks_country_id` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
;