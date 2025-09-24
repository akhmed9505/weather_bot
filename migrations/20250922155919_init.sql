-- +goose Up
-- +goose StatementBegin
CREATE TABLE users (
    id BIGINT primary key,
    city text,
    created_at timestamp default NOW()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users;
-- +goose StatementEnd
