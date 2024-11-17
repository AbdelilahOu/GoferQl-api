package resolvers

import (
	"fmt"

	"github.com/AbdelilahOu/GoferQl/graphql/utils"
	db "github.com/AbdelilahOu/GoferQl/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

func ListPosts(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	var Limit int32 = 20
	var Offset int32 = 0

	if val, ok := p.Args["limit"]; ok && val != nil {
		Limit = val.(int32)
	}
	if val, ok := p.Args["offset"]; ok && val != nil {
		Offset = val.(int32)
	}

	return dbQueries.ListPosts(p.Context, db.ListPostsParams{
		Limit:  Limit,
		Offset: Offset,
	})
}

func ListPostsByUserID(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	var Limit int32 = 20
	var Offset int32 = 0
	var UserID uuid.UUID

	if val, ok := p.Args["limit"]; ok && val != nil {
		Limit = val.(int32)
	}
	if val, ok := p.Args["offset"]; ok && val != nil {
		Offset = val.(int32)
	}
	if val, ok := p.Source.(db.User); ok {
		UserID = val.ID
	}

	return dbQueries.ListPostsByUserID(p.Context, db.ListPostsByUserIDParams{
		UserID: utils.UuidToPgTypeUuid(UserID),
		Limit:  Limit,
		Offset: Offset,
	})
}

func GetPost(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	if val, ok := p.Args["id"]; ok && val != nil {
		ID, err := uuid.Parse(val.(string))
		if err != nil {
			return nil, err
		}
		return dbQueries.GetPost(p.Context, ID)
	}

	return nil, fmt.Errorf("id not provided")
}

func CreatePost(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	UserID, err := uuid.Parse(p.Args["userId"].(string))
	if err != nil {
		return nil, err
	}

	CategoryID, err := uuid.Parse(p.Args["categoryId"].(string))
	if err != nil {
		return nil, err
	}

	params := db.CreatePostParams{
		Title:      p.Args["title"].(string),
		Content:    p.Args["content"].(string),
		Status:     utils.NullablePgTypeText(p.Args, "status"),
		UserID:     utils.UuidToPgTypeUuid(UserID),
		CategoryID: utils.UuidToPgTypeUuid(CategoryID),
	}

	return dbQueries.CreatePost(p.Context, params)
}

func UpdatePost(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	ID, err := uuid.Parse(p.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	CategoryID, err := uuid.Parse(p.Args["categoryId"].(string))
	if err != nil {
		return nil, err
	}

	params := db.UpdatePostParams{
		ID:         ID,
		Title:      utils.NullablePgTypeText(p.Args, "title"),
		Content:    utils.NullablePgTypeText(p.Args, "content"),
		Status:     utils.NullablePgTypeText(p.Args, "status"),
		CategoryID: utils.UuidToPgTypeUuid(CategoryID),
	}

	return dbQueries.UpdatePost(p.Context, params)
}

func DeletePost(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)
	id := p.Args["id"].(string)

	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return dbQueries.DeletePost(p.Context, ID)
}
