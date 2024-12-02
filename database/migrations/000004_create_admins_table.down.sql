ALTER TABLE `admins`
    DROP INDEX `index_username`,
    DROP INDEX `index_email`;

DROP TABLE IF EXISTS `admins`;