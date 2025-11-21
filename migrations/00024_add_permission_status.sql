-- +goose Up
-- +goose StatementBegin
ALTER TABLE `permissions` 
ADD COLUMN `permission_status` ENUM('main_menu', 'submenu', 'action') NOT NULL DEFAULT 'action' AFTER `permission_description`;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE `permissions` 
DROP COLUMN `permission_status`;
-- +goose StatementEnd
