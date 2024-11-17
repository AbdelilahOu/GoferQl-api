package resolvers

import (
	"fmt"

	"github.com/AbdelilahOu/GoferQl/graphql/utils"
	db "github.com/AbdelilahOu/GoferQl/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

func ListUsers(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	var Limit int32 = 20
	var Offset int32 = 0

	if val, ok := p.Args["limit"]; ok && val != nil {
		Limit = val.(int32)
	}
	if val, ok := p.Args["offset"]; ok && val != nil {
		Offset = val.(int32)
	}

	return dbQueries.ListUsers(p.Context, db.ListUsersParams{
		Limit:  Limit,
		Offset: Offset,
	})
}

func GetUser(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	if val, ok := p.Args["id"]; ok && val != nil {
		ID, err := uuid.Parse(val.(string))
		if err != nil {
			return nil, err
		}
		return dbQueries.GetUser(p.Context, ID)
	}
	if val, ok := p.Args["email"]; ok && val != nil {
		return dbQueries.GetUserByEmail(p.Context, val.(string))
	}

	return nil, fmt.Errorf("id or email not provided")
}

func GetPostUser(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	var PostID uuid.UUID

	if val, ok := p.Source.(db.Post); ok {
		PostID = val.ID
	}

	return dbQueries.GetUserByPostID(p.Context, PostID)
}

func GetCommentUser(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	var CommentID uuid.UUID

	if val, ok := p.Source.(db.Comment); ok {
		CommentID = val.ID
	}

	return dbQueries.GetUserByCommentID(p.Context, CommentID)
}

func CreateUser(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	params := db.CreateUserParams{
		Username: p.Args["username"].(string),
		Email:    p.Args["email"].(string),
		Password: p.Args["password"].(string),
		Bio:      utils.NullablePgTypeText(p.Args, "bio"),
	}

	return dbQueries.CreateUser(p.Context, params)
}

func UpdateUser(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	ID, err := uuid.Parse(p.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	params := db.UpdateUserParams{
		ID:       ID,
		Username: utils.NullablePgTypeText(p.Args, "username"),
		Email:    utils.NullablePgTypeText(p.Args, "email"),
		Password: utils.NullablePgTypeText(p.Args, "password"),
		Bio:      utils.NullablePgTypeText(p.Args, "bio"),
	}

	return dbQueries.UpdateUser(p.Context, params)
}

func DeleteUser(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)
	id := p.Args["id"].(string)

	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return dbQueries.DeleteUser(p.Context, ID)
}
