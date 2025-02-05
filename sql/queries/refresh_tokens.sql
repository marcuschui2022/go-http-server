-- name: CreateRefreshToken :one
insert into refresh_tokens(token, created_at, updated_at, user_id, expires_at, revoked_at)
values ($1, now(), now(), $2, now() + interval '60 days', null)
returning *;

-- name: GetUserFromRefreshToken :one
select u.*
from refresh_tokens rt
         join users u
              on rt.user_id = u.id
where token = $1
  and revoked_at is null
  and expires_at > now();

-- name: RevokeToken :exec
update refresh_tokens
set revoked_at = now(),
    updated_at=now()
where token = $1;