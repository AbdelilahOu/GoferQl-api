
-- name: CreatePost :one
INSERT INTO posts (
    title,
    content,
    user_id,
    category_id,
    status
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING *;

-- name: GetPost :one
SELECT 
    p.*,
    u.username as author_name,
    c.name as category_name
FROM posts p
LEFT JOIN users u ON p.user_id = u.id
LEFT JOIN categories c ON p.category_id = c.id
WHERE p.id = $1;

-- name: UpdatePost :one
UPDATE posts
SET 
    title = COALESCE($2, title),
    content = COALESCE($3, content),
    category_id = COALESCE($4, category_id),
    status = COALESCE($5, status),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1 AND user_id = $6
RETURNING *;

-- name: ListPosts :many
SELECT 
    p.*,
    u.username as author_name,
    c.name as category_name
FROM posts p
LEFT JOIN users u ON p.user_id = u.id
LEFT JOIN categories c ON p.category_id = c.id
WHERE 
    ($1::varchar IS NULL OR p.status = $1) AND
    ($2::integer IS NULL OR p.category_id = $2) AND
    ($3::integer IS NULL OR p.user_id = $3)
ORDER BY p.created_at DESC
LIMIT $4 OFFSET $5;

-- name: DeletePost :exec
DELETE FROM posts
WHERE id = $1 AND user_id = $2;

