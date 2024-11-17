package resolvers

import (
	db "github.com/AbdelilahOu/GoferQl/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

func ListPostTags(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	var PostID uuid.UUID
	if val, ok := p.Source.(db.Post); ok {
		PostID = val.ID
	}

	return dbQueries.ListPostTags(p.Context, PostID)
}

func AddPostTag(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	PostID, err := uuid.Parse(p.Args["postId"].(string))
	if err != nil {
		return nil, err
	}

	TagID, err := uuid.Parse(p.Args["tagId"].(string))
	if err != nil {
		return nil, err
	}

	params := db.AddPostTagParams{
		PostID: PostID,
		TagID:  TagID,
	}

	err = dbQueries.AddPostTag(p.Context, params)
	if err != nil {
		return nil, err
	}

	return 1, nil
}

func RemovePostTag(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	PostID, err := uuid.Parse(p.Args["postId"].(string))
	if err != nil {
		return nil, err
	}

	TagID, err := uuid.Parse(p.Args["tagId"].(string))
	if err != nil {
		return nil, err
	}

	params := db.AddPostTagParams{
		PostID: PostID,
		TagID:  TagID,
	}

	err = dbQueries.RemovePostTag(p.Context, db.RemovePostTagParams(params))
	if err != nil {
		return nil, err
	}

	return 1, nil
}
