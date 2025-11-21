-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `permissions` (
  `permission_id` INT NOT NULL AUTO_INCREMENT,
  `permission_name` VARCHAR(100) NOT NULL UNIQUE,
  `permission_resource` VARCHAR(50) NOT NULL,
  `permission_action` VARCHAR(50) NOT NULL,
  `permission_description` TEXT,
  `permission_create_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  `permission_update_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`permission_id`),
  INDEX `idx_resource_action` (`permission_resource`, `permission_action`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS `permissions`;
-- +goose StatementEnd
