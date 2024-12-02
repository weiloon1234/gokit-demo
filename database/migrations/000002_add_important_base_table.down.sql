ALTER TABLE `banks`
DROP INDEX `index_country_id`,
DROP FOREIGN KEY `fk_banks_country_id`;

-- Drop the table
DROP TABLE IF EXISTS `country_locations`;
DROP TABLE IF EXISTS `banks`;
