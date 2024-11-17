-- name: CreateUser :one
INSERT INTO users (
    username,
    email,
    password,
    bio
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: UpdateUser :one
UPDATE users
SET 
    username = COALESCE(sqlc.narg(username), username),
    email = COALESCE(sqlc.narg(email), email),
    password = COALESCE(sqlc.narg(password), password),
    bio = COALESCE(sqlc.narg(bio), bio)
WHERE id = $1
RETURNING id, username, email, bio;

-- name: DeleteUser :one
DELETE FROM users
WHERE id = $1 RETURNING id;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;