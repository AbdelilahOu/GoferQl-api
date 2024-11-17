// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	AddPostTag(ctx context.Context, arg AddPostTagParams) error
	CreateCategory(ctx context.Context, arg CreateCategoryParams) (Category, error)
	CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error)
	CreatePost(ctx context.Context, arg CreatePostParams) (Post, error)
	CreateTag(ctx context.Context, name string) (Tag, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteCategory(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	DeleteComment(ctx context.Context, arg DeleteCommentParams) error
	DeletePost(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	DeleteUser(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	GetCategory(ctx context.Context, id uuid.UUID) (Category, error)
	GetComment(ctx context.Context, id uuid.UUID) (GetCommentRow, error)
	GetPost(ctx context.Context, id uuid.UUID) (GetPostRow, error)
	GetTag(ctx context.Context, id uuid.UUID) (Tag, error)
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
	GetUserByEmail(ctx context.Context, email string) (User, error)
	ListCategories(ctx context.Context, arg ListCategoriesParams) ([]Category, error)
	ListCommentsByPost(ctx context.Context, postID pgtype.UUID) ([]ListCommentsByPostRow, error)
	ListPostTags(ctx context.Context, postID uuid.UUID) ([]Tag, error)
	ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error)
	ListPostsByTag(ctx context.Context, arg ListPostsByTagParams) ([]ListPostsByTagRow, error)
	ListTags(ctx context.Context) ([]Tag, error)
	ListUsers(ctx context.Context, arg ListUsersParams) ([]User, error)
	RemovePostTag(ctx context.Context, arg RemovePostTagParams) error
	UpdateCategory(ctx context.Context, arg UpdateCategoryParams) (Category, error)
	UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (UpdateUserRow, error)
}

var _ Querier = (*Queries)(nil)