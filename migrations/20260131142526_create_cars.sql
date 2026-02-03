-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS car (
    car_id SERIAL PRIMARY KEY,
    mark VARCHAR(255),
    model VARCHAR(255),
    owner_count INT,
    price INT,
    currency VARCHAR(3),
    options VARCHAR(255)[]
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS car;
-- +goose StatementEnd