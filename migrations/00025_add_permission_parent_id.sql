-- +goose Up
-- +goose StatementBegin
ALTER TABLE `permissions` 
ADD COLUMN `permission_parent_id` INT NULL AFTER `permission_status`,
ADD CONSTRAINT `fk_permission_parent` FOREIGN KEY (`permission_parent_id`) REFERENCES `permissions` (`permission_id`) ON DELETE CASCADE,
ADD INDEX `idx_parent_id` (`permission_parent_id`);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `permissions` 
DROP FOREIGN KEY `fk_permission_parent`,
DROP INDEX `idx_parent_id`,
DROP COLUMN `permission_parent_id`;
-- +goose StatementEnd
