// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: comments.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createComment = `-- name: CreateComment :one
INSERT INTO comments (
    content,
    post_id,
    user_id,
    parent_id
) VALUES (
    $1, $2, $3, $4
) RETURNING id, content, post_id, user_id, parent_id, created_at
`

type CreateCommentParams struct {
	Content  string      `json:"content"`
	PostID   pgtype.UUID `json:"post_id"`
	UserID   pgtype.UUID `json:"user_id"`
	ParentID pgtype.UUID `json:"parent_id"`
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, createComment,
		arg.Content,
		arg.PostID,
		arg.UserID,
		arg.ParentID,
	)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.PostID,
		&i.UserID,
		&i.ParentID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteComment = `-- name: DeleteComment :one
DELETE FROM comments
WHERE id = $1 RETURNING id
`

func (q *Queries) DeleteComment(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, deleteComment, id)
	err := row.Scan(&id)
	return id, err
}

const listCommentsByParentID = `-- name: ListCommentsByParentID :many
SELECT 
    id, content, post_id, user_id, parent_id, created_at
FROM comments
WHERE parent_id = $1
ORDER BY created_at DESC
`

func (q *Queries) ListCommentsByParentID(ctx context.Context, parentID pgtype.UUID) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsByParentID, parentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.PostID,
			&i.UserID,
			&i.ParentID,
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

const listCommentsByPostID = `-- name: ListCommentsByPostID :many
SELECT 
    id, content, post_id, user_id, parent_id, created_at
FROM comments
WHERE post_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListCommentsByPostIDParams struct {
	PostID pgtype.UUID `json:"post_id"`
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
}

func (q *Queries) ListCommentsByPostID(ctx context.Context, arg ListCommentsByPostIDParams) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsByPostID, arg.PostID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.PostID,
			&i.UserID,
			&i.ParentID,
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

const listCommentsByUserID = `-- name: ListCommentsByUserID :many
SELECT 
    id, content, post_id, user_id, parent_id, created_at
FROM comments
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListCommentsByUserIDParams struct {
	UserID pgtype.UUID `json:"user_id"`
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
}

func (q *Queries) ListCommentsByUserID(ctx context.Context, arg ListCommentsByUserIDParams) ([]Comment, error) {
	rows, err := q.db.Query(ctx, listCommentsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Comment{}
	for rows.Next() {
		var i Comment
		if err := rows.Scan(
			&i.ID,
			&i.Content,
			&i.PostID,
			&i.UserID,
			&i.ParentID,
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

const updateComment = `-- name: UpdateComment :one
UPDATE comments
SET 
    content = $2
WHERE id = $1
RETURNING id, content, post_id, user_id, parent_id, created_at
`

type UpdateCommentParams struct {
	ID      uuid.UUID `json:"id"`
	Content string    `json:"content"`
}

func (q *Queries) UpdateComment(ctx context.Context, arg UpdateCommentParams) (Comment, error) {
	row := q.db.QueryRow(ctx, updateComment, arg.ID, arg.Content)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Content,
		&i.PostID,
		&i.UserID,
		&i.ParentID,
		&i.CreatedAt,
	)
	return i, err
}
