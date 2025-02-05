// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: refresh_tokens.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createRefreshToken = `-- name: CreateRefreshToken :one
insert into refresh_tokens(token, created_at, updated_at, user_id, expires_at, revoked_at)
values ($1, now(), now(), $2, now() + interval '60 days', null)
returning token, created_at, updated_at, expires_at, revoked_at, user_id
`

type CreateRefreshTokenParams struct {
	Token  string
	UserID uuid.UUID
}

func (q *Queries) CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (RefreshToken, error) {
	row := q.db.QueryRowContext(ctx, createRefreshToken, arg.Token, arg.UserID)
	var i RefreshToken
	err := row.Scan(
		&i.Token,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.ExpiresAt,
		&i.RevokedAt,
		&i.UserID,
	)
	return i, err
}

const getUserFromRefreshToken = `-- name: GetUserFromRefreshToken :one
select u.id, u.created_at, u.updated_at, u.email, u.hashed_password, u.is_chirpy_red
from refresh_tokens rt
         join users u
              on rt.user_id = u.id
where token = $1
  and revoked_at is null
  and expires_at > now()
`

func (q *Queries) GetUserFromRefreshToken(ctx context.Context, token string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserFromRefreshToken, token)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Email,
		&i.HashedPassword,
		&i.IsChirpyRed,
	)
	return i, err
}

const revokeToken = `-- name: RevokeToken :exec
update refresh_tokens
set revoked_at = now(),
    updated_at=now()
where token = $1
`

func (q *Queries) RevokeToken(ctx context.Context, token string) error {
	_, err := q.db.ExecContext(ctx, revokeToken, token)
	return err
}
