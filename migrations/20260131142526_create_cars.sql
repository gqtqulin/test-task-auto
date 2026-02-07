-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS car (
    car_id BIGSERIAL PRIMARY KEY,
    mark VARCHAR(255),
    model VARCHAR(255),
    owner_count BIGINT,
    price BIGINT,
    currency VARCHAR(3),
    options VARCHAR(255)[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS car;
-- +goose StatementEnd