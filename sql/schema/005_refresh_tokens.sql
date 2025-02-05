-- +goose Up
alter table refresh_tokens
    alter column user_id set not null,
    alter column revoked_at drop not null;


-- +goose Down
alter table refresh_tokens
    alter column user_id drop not null,
    alter column revoked_at set not null;