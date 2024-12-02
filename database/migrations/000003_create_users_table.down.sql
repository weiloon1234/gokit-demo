-- Drop foreign key constraints
ALTER TABLE `users`
DROP FOREIGN KEY `fk_users_country_id`,
DROP FOREIGN KEY `fk_users_contact_country_id`,
DROP FOREIGN KEY `fk_users_introducer_user_id`,
DROP FOREIGN KEY `fk_users_bank_id`;

-- Drop indexes
ALTER TABLE `users`
DROP INDEX `index_username`,
DROP INDEX `index_email`,
DROP INDEX `fk_country_id`,
DROP INDEX `fk_contact_country_id`,
DROP INDEX `fk_introducer_user_id`,
DROP INDEX `fk_bank_id`;

-- Drop the table
DROP TABLE IF EXISTS `users`;
