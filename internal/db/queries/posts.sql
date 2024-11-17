
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
    title = COALESCE(sqlc.narg(title), title),
    content = COALESCE(sqlc.narg(content), content),
    category_id = COALESCE(sqlc.narg(category_id), category_id),
    status = COALESCE(sqlc.narg(status), status),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING *;

-- name: ListPosts :many
SELECT 
    *
FROM posts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2;

-- name: ListPostsByUserID :many
SELECT 
    *
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListPostsByTagID :many
SELECT 
    p.*
FROM posts as p
JOIN post_tags as ps ON ps.post_id = p.id
WHERE ps.tag_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;


-- name: DeletePost :one
DELETE FROM posts
WHERE id = $1 RETURNING id;

