-- +goose Up
-- +goose StatementBegin
INSERT INTO car (mark, model, owner_count, price, currency, options) VALUES
('Toyota', 'RAV4', 1, 3000000, 'RUB', ARRAY['4wd', 'something_else']),
('Mercedes', 'C-Class', 1, 4500000, 'RUB', ARRAY['premium_audio', 'something_else']),
('Tesla', 'Model 3', 2, 5000000, 'USD', ARRAY['something', 'something_2'])
-- +goose StatementEnd


