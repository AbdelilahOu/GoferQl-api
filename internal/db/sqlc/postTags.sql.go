// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: postTags.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const addPostTag = `-- name: AddPostTag :exec
INSERT INTO post_tags (post_id, tag_id)
VALUES ($1, $2)
`

type AddPostTagParams struct {
	PostID uuid.UUID `json:"post_id"`
	TagID  uuid.UUID `json:"tag_id"`
}

func (q *Queries) AddPostTag(ctx context.Context, arg AddPostTagParams) error {
	_, err := q.db.Exec(ctx, addPostTag, arg.PostID, arg.TagID)
	return err
}

const listPostTags = `-- name: ListPostTags :many
SELECT t.id, t.name
FROM tags t
JOIN post_tags pt ON pt.tag_id = t.id
WHERE pt.post_id = $1
ORDER BY t.name
`

func (q *Queries) ListPostTags(ctx context.Context, postID uuid.UUID) ([]Tag, error) {
	rows, err := q.db.Query(ctx, listPostTags, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Tag{}
	for rows.Next() {
		var i Tag
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listPostsByTag = `-- name: ListPostsByTag :many
SELECT 
    p.id, p.title, p.content, p.user_id, p.category_id, p.status, p.created_at, p.updated_at,
    u.username as author_name,
    c.name as category_name
FROM posts p
JOIN post_tags pt ON pt.post_id = p.id
LEFT JOIN users u ON p.user_id = u.id
LEFT JOIN categories c ON p.category_id = c.id
WHERE pt.tag_id = $1
ORDER BY p.created_at DESC
LIMIT $2 OFFSET $3
`

type ListPostsByTagParams struct {
	TagID  uuid.UUID `json:"tag_id"`
	Limit  int32     `json:"limit"`
	Offset int32     `json:"offset"`
}

type ListPostsByTagRow struct {
	ID           uuid.UUID          `json:"id"`
	Title        string             `json:"title"`
	Content      string             `json:"content"`
	UserID       pgtype.UUID        `json:"user_id"`
	CategoryID   pgtype.UUID        `json:"category_id"`
	Status       pgtype.Text        `json:"status"`
	CreatedAt    pgtype.Timestamptz `json:"created_at"`
	UpdatedAt    pgtype.Timestamptz `json:"updated_at"`
	AuthorName   pgtype.Text        `json:"author_name"`
	CategoryName pgtype.Text        `json:"category_name"`
}

func (q *Queries) ListPostsByTag(ctx context.Context, arg ListPostsByTagParams) ([]ListPostsByTagRow, error) {
	rows, err := q.db.Query(ctx, listPostsByTag, arg.TagID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListPostsByTagRow{}
	for rows.Next() {
		var i ListPostsByTagRow
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.UserID,
			&i.CategoryID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.AuthorName,
			&i.CategoryName,
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

const removePostTag = `-- name: RemovePostTag :exec
DELETE FROM post_tags
WHERE post_id = $1 AND tag_id = $2
`

type RemovePostTagParams struct {
	PostID uuid.UUID `json:"post_id"`
	TagID  uuid.UUID `json:"tag_id"`
}

func (q *Queries) RemovePostTag(ctx context.Context, arg RemovePostTagParams) error {
	_, err := q.db.Exec(ctx, removePostTag, arg.PostID, arg.TagID)
	return err
}
