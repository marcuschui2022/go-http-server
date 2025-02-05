-- +goose Up
alter table refresh_tokens
    drop column user_id,
    add column user_id uuid not null references users (id) on delete cascade;


-- +goose Down
alter table refresh_tokens
    drop column user_id,
    add column user_id uuid not null references users (id);