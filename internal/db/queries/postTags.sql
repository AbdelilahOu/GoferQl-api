-- name: AddPostTag :exec
INSERT INTO post_tags (post_id, tag_id)
VALUES ($1, $2);

-- name: RemovePostTag :exec
DELETE FROM post_tags
WHERE post_id = $1 AND tag_id = $2;

-- name: ListPostTags :many
SELECT t.*
FROM tags t
JOIN post_tags pt ON pt.tag_id = t.id
WHERE pt.post_id = $1
ORDER BY t.name;

-- name: ListPostsByTag :many
SELECT 
    p.*,
    u.username as author_name,
    c.name as category_name
FROM posts p
JOIN post_tags pt ON pt.post_id = p.id
LEFT JOIN users u ON p.user_id = u.id
LEFT JOIN categories c ON p.category_id = c.id
WHERE pt.tag_id = $1
ORDER BY p.created_at DESC
LIMIT $2 OFFSET $3;