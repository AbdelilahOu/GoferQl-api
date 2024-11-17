
-- name: CreateCategory :one
INSERT INTO categories (
    name,
    description
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetCategory :one
SELECT * FROM categories
WHERE id = $1;

-- name: UpdateCategory :one
UPDATE categories
SET 
    name = COALESCE(sqlc.narg(name), name),
    description = COALESCE(sqlc.narg(description), description)
WHERE id = $1
RETURNING *;

-- name: ListCategories :many
SELECT * FROM categories
ORDER BY name
LIMIT $1 OFFSET $2;

-- name: DeleteCategory :one
DELETE FROM categories
WHERE id = $1 RETURNING id;
