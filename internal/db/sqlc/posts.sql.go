// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: posts.sql

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
    title,
    content,
    user_id,
    category_id,
    status
) VALUES (
    $1, $2, $3, $4, $5
) RETURNING id, title, content, user_id, category_id, status, created_at, updated_at
`

type CreatePostParams struct {
	Title      string      `json:"title"`
	Content    string      `json:"content"`
	UserID     pgtype.UUID `json:"user_id"`
	CategoryID pgtype.UUID `json:"category_id"`
	Status     pgtype.Text `json:"status"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, createPost,
		arg.Title,
		arg.Content,
		arg.UserID,
		arg.CategoryID,
		arg.Status,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.UserID,
		&i.CategoryID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :one
DELETE FROM posts
WHERE id = $1 RETURNING id
`

func (q *Queries) DeletePost(ctx context.Context, id uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, deletePost, id)
	err := row.Scan(&id)
	return id, err
}

const getPost = `-- name: GetPost :one
SELECT 
    p.id, p.title, p.content, p.user_id, p.category_id, p.status, p.created_at, p.updated_at,
    u.username as author_name,
    c.name as category_name
FROM posts p
LEFT JOIN users u ON p.user_id = u.id
LEFT JOIN categories c ON p.category_id = c.id
WHERE p.id = $1
`

type GetPostRow struct {
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

func (q *Queries) GetPost(ctx context.Context, id uuid.UUID) (GetPostRow, error) {
	row := q.db.QueryRow(ctx, getPost, id)
	var i GetPostRow
	err := row.Scan(
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
	)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT 
    id, title, content, user_id, category_id, status, created_at, updated_at
FROM posts
ORDER BY created_at DESC
LIMIT $1 OFFSET $2
`

type ListPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error) {
	rows, err := q.db.Query(ctx, listPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.UserID,
			&i.CategoryID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const listPostsByUserID = `-- name: ListPostsByUserID :many
SELECT 
    id, title, content, user_id, category_id, status, created_at, updated_at
FROM posts
WHERE user_id = $1
ORDER BY created_at DESC
LIMIT $2 OFFSET $3
`

type ListPostsByUserIDParams struct {
	UserID pgtype.UUID `json:"user_id"`
	Limit  int32       `json:"limit"`
	Offset int32       `json:"offset"`
}

func (q *Queries) ListPostsByUserID(ctx context.Context, arg ListPostsByUserIDParams) ([]Post, error) {
	rows, err := q.db.Query(ctx, listPostsByUserID, arg.UserID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Content,
			&i.UserID,
			&i.CategoryID,
			&i.Status,
			&i.CreatedAt,
			&i.UpdatedAt,
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

const updatePost = `-- name: UpdatePost :one
UPDATE posts
SET 
    title = COALESCE($2, title),
    content = COALESCE($3, content),
    category_id = COALESCE($4, category_id),
    status = COALESCE($5, status),
    updated_at = CURRENT_TIMESTAMP
WHERE id = $1
RETURNING id, title, content, user_id, category_id, status, created_at, updated_at
`

type UpdatePostParams struct {
	ID         uuid.UUID   `json:"id"`
	Title      pgtype.Text `json:"title"`
	Content    pgtype.Text `json:"content"`
	CategoryID pgtype.UUID `json:"category_id"`
	Status     pgtype.Text `json:"status"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.db.QueryRow(ctx, updatePost,
		arg.ID,
		arg.Title,
		arg.Content,
		arg.CategoryID,
		arg.Status,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Content,
		&i.UserID,
		&i.CategoryID,
		&i.Status,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
