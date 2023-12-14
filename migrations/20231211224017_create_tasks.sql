-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS "tasks" (
     id uuid,
     title varchar(255),
     description text DEFAULT null,
     completed bool default false,
     create_at timestamp not null,
     update_at timestamp not null,
     delete_at timestamp null
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "tasks";
-- +goose StatementEnd
