ALTER TABLE `country_locations`
DROP INDEX `index_country_id`,
DROP INDEX `index_parent_id`,

DROP FOREIGN KEY `fk_country_locations_country_id`,
DROP FOREIGN KEY `fk_country_locations_parent_id`;

ALTER TABLE `countries`
DROP INDEX `index_iso2`,
DROP INDEX `index_iso3`,
DROP INDEX `index_status`;

DROP TABLE IF EXISTS `country_locations`;
DROP TABLE IF EXISTS `countries`;
