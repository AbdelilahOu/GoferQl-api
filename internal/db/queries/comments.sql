-- name: CreateComment :one
INSERT INTO comments (
    content,
    post_id,
    user_id,
    parent_id
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: UpdateComment :one
UPDATE comments
SET 
    content = $2
WHERE id = $1
RETURNING *;

-- name: ListCommentsByPostID :many
SELECT 
    *
FROM comments
WHERE post_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: ListCommentsByParentID :many
SELECT 
    *
FROM comments
WHERE parent_id = $1
ORDER BY created_at DESC;

-- name: ListCommentsByUserID :many
SELECT 
    *
FROM comments
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3;

-- name: DeleteComment :one
DELETE FROM comments
WHERE id = $1 RETURNING id;