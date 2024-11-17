-- name: CreateComment :one
INSERT INTO comments (
    content,
    post_id,
    user_id,
    parent_id
) VALUES (
    $1, $2, $3, $4
) RETURNING *;

-- name: GetComment :one
SELECT 
    c.*,
    u.username as author_name
FROM comments c
LEFT JOIN users u ON c.user_id = u.id
WHERE c.id = $1;

-- name: ListCommentsByPost :many
SELECT 
    c.*,
    u.username as author_name
FROM comments c
LEFT JOIN users u ON c.user_id = u.id
WHERE c.post_id = $1
ORDER BY c.created_at DESC;

-- name: DeleteComment :exec
DELETE FROM comments
WHERE id = $1 AND user_id = $2;