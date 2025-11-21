-- +goose Up
-- +goose StatementBegin
-- View will be created manually or by application
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- DROP VIEW IF EXISTS v_order_detail;
-- +goose StatementEnd
