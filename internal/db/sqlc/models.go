// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Category struct {
	ID          uuid.UUID          `json:"id"`
	Name        string             `json:"name"`
	Description pgtype.Text        `json:"description"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
}

type Comment struct {
	ID        uuid.UUID          `json:"id"`
	Content   string             `json:"content"`
	PostID    pgtype.UUID        `json:"post_id"`
	UserID    pgtype.UUID        `json:"user_id"`
	ParentID  pgtype.UUID        `json:"parent_id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type Post struct {
	ID         uuid.UUID          `json:"id"`
	Title      string             `json:"title"`
	Content    string             `json:"content"`
	UserID     pgtype.UUID        `json:"user_id"`
	CategoryID pgtype.UUID        `json:"category_id"`
	Status     pgtype.Text        `json:"status"`
	CreatedAt  pgtype.Timestamptz `json:"created_at"`
	UpdatedAt  pgtype.Timestamptz `json:"updated_at"`
}

type PostTag struct {
	PostID uuid.UUID `json:"post_id"`
	TagID  uuid.UUID `json:"tag_id"`
}

type Tag struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

type User struct {
	ID        uuid.UUID          `json:"id"`
	Username  string             `json:"username"`
	Email     string             `json:"email"`
	Password  string             `json:"password"`
	Bio       pgtype.Text        `json:"bio"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}
