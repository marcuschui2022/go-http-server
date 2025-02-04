-- +goose Up
alter table users
    add column hashed_password TEXT not null default 'unset';

-- +goose Down
alter table users
drop
column hashed_password;