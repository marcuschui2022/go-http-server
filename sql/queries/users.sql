-- name: CreateUser :one
insert into users (id, created_at, updated_at, email, hashed_password)
values (gen_random_uuid(), now(), now(), $1, $2)
returning *;

-- name: GetUserByEmail :one
select *
from users
where email = $1;

-- name: UpdateUser :one
update users
set email=$2,
    hashed_password =$3,
    updated_at=now()
where id = $1
returning *;

-- name: UpgradeToChirpyRed :one
UPDATE users
SET is_chirpy_red = true,
    updated_at    = NOW()
WHERE id = $1
RETURNING *;