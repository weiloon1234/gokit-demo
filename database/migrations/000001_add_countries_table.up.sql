CREATE TABLE `countries` (
    `id` bigint UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `iso2` CHAR(2) NOT NULL COMMENT 'ISO 3166-1 alpha-2 country code',
    `iso3` CHAR(3) NOT NULL COMMENT 'ISO 3166-1 alpha-3 country code',
    `name` VARCHAR(255) NOT NULL COMMENT 'Country name',
    `official_name` VARCHAR(255) DEFAULT NULL COMMENT 'Official country name',
    `numeric_code` CHAR(3) DEFAULT NULL COMMENT 'ISO 3166-1 numeric code',
    `phone_code` VARCHAR(10) DEFAULT NULL COMMENT 'Country calling code, e.g. +1 for US',
    `capital` VARCHAR(255) DEFAULT NULL COMMENT 'Capital city of the country',
    `currency_name` VARCHAR(255) DEFAULT NULL COMMENT 'Currency name',
    `currency_code` CHAR(3) DEFAULT NULL COMMENT 'ISO 4217 currency code',
    `currency_symbol` VARCHAR(255) DEFAULT NULL COMMENT 'Currency symbol',
    `conversion_rate` decimal(18,5) NOT NULL DEFAULT '0.00000',
    `status` tinyint UNSIGNED NOT NULL DEFAULT '0',
    `created_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP COMMENT 'Record creation timestamp',
    `updated_at` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE `countries`
    ADD INDEX `index_iso2` (`iso2`),
    ADD INDEX `index_iso3` (`iso3`),
    ADD INDEX `index_status` (`status`);

CREATE TABLE `country_locations` (
    `id` bigint UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `country_id` bigint UNSIGNED DEFAULT NULL COMMENT 'Country of the location',
    `parent_id` bigint UNSIGNED DEFAULT NULL COMMENT 'If present, it is area, this is state id',
    `sorting` bigint UNSIGNED DEFAULT '0',
    `name_en` VARCHAR(255) NOT NULL COMMENT 'Bank name in english',
    `name_zh` VARCHAR(255) NOT NULL COMMENT 'Bank name in chinese',
    `created_at` timestamp NULL DEFAULT NULL COMMENT 'Record creation timestamp',
    `updated_at` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT 'Record update timestamp',
    `deleted_at` timestamp NULL DEFAULT NULL COMMENT 'Record deleted timestamp'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE `country_locations`
    ADD INDEX `index_country_id` (`country_id`),
    ADD INDEX `index_parent_id` (`parent_id`),

    ADD CONSTRAINT `fk_country_locations_country_id` FOREIGN KEY (`country_id`) REFERENCES `countries` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
    ADD CONSTRAINT `fk_country_locations_parent_id` FOREIGN KEY (`parent_id`) REFERENCES `country_locations` (`id`) ON DELETE SET NULL ON UPDATE CASCADE
;