// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password,
    bio
) VALUES (
    $1, $2, $3, $4
) RETURNING id, username, email, password, bio, created_at
`

type CreateUserParams struct {
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Password string      `json:"password"`
	Bio      pgtype.Text `json:"bio"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Bio,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.CreatedAt,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1 RETURNING id
`

func (q *Queries) DeleteUser(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, deleteUser, id)
	err := row.Scan(&id)
	return id, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, email, password, bio, created_at FROM users
WHERE id = $1
`

func (q *Queries) GetUser(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.CreatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, username, email, password, bio, created_at FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Password,
		&i.Bio,
		&i.CreatedAt,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
SELECT id, username, email, password, bio, created_at FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type ListUsersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error) {
	rows, err := q.db.Query(ctx, listUsers, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Username,
			&i.Email,
			&i.Password,
			&i.Bio,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateUser = `-- name: UpdateUser :one
UPDATE users
SET 
    username = COALESCE($2, username),
    email = COALESCE($3, email),
    password = COALESCE($4, password),
    bio = COALESCE($5, bio)
WHERE id = $1
RETURNING id, username, email, bio
`

type UpdateUserParams struct {
	ID       uuid.UUID   `json:"id"`
	Username pgtype.Text `json:"username"`
	Email    pgtype.Text `json:"email"`
	Password pgtype.Text `json:"password"`
	Bio      pgtype.Text `json:"bio"`
}

type UpdateUserRow struct {
	ID       uuid.UUID   `json:"id"`
	Username string      `json:"username"`
	Email    string      `json:"email"`
	Bio      pgtype.Text `json:"bio"`
}

func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error) {
	row := q.db.QueryRow(ctx, updateUser,
		arg.ID,
		arg.Username,
		arg.Email,
		arg.Password,
		arg.Bio,
	)
	var i UpdateUserRow
	err := row.Scan(
		&i.ID,
		&i.Username,
		&i.Email,
		&i.Bio,
	)
	return i, err
}
