package resolvers

import (
	"github.com/AbdelilahOu/GoferQl/graphql/utils"
	db "github.com/AbdelilahOu/GoferQl/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

func ListCommentsByPostID(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	var Limit int = 20
	var Offset int = 0

	ID, err := uuid.Parse(p.Args["postId"].(string))
	if err != nil {
		return nil, err
	}
	if val, ok := p.Args["limit"]; ok && val != nil {
		Limit = val.(int)
	}
	if val, ok := p.Args["offset"]; ok && val != nil {
		Offset = val.(int)
	}

	return dbQueries.ListCommentsByPostID(p.Context, db.ListCommentsByPostIDParams{
		PostID: utils.UuidToPgTypeUuid(ID),
		Limit:  int32(Limit),
		Offset: int32(Offset),
	})
}

func ListPostComments(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	var Limit int = 20
	var PostID uuid.UUID

	if val, ok := p.Args["commentsLimit"]; ok && val != nil {
		Limit = val.(int)
	}
	if val, ok := p.Source.(db.Post); ok {
		PostID = val.ID
	}

	return dbQueries.ListCommentsByPostID(p.Context, db.ListCommentsByPostIDParams{
		PostID: utils.UuidToPgTypeUuid(PostID),
		Limit:  int32(Limit),
		Offset: 0,
	})
}

func ListCommentChildren(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	var CommentID uuid.UUID

	if val, ok := p.Source.(db.Comment); ok {
		CommentID = val.ID
	}

	return dbQueries.ListCommentsByParentID(p.Context, utils.UuidToPgTypeUuid(CommentID))
}

func CreateComment(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	UserID, err := uuid.Parse(p.Args["userId"].(string))
	if err != nil {
		return nil, err
	}

	PostID, err := uuid.Parse(p.Args["postId"].(string))
	if err != nil {
		return nil, err
	}

	ParentID, err := uuid.Parse(p.Args["parentId"].(string))
	if err != nil {
		return nil, err
	}

	params := db.CreateCommentParams{
		Content:  p.Args["content"].(string),
		UserID:   utils.UuidToPgTypeUuid(UserID),
		PostID:   utils.UuidToPgTypeUuid(PostID),
		ParentID: utils.UuidToPgTypeUuid(ParentID),
	}

	return dbQueries.CreateComment(p.Context, params)
}

func UpdateComment(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	ID, err := uuid.Parse(p.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	params := db.UpdateCommentParams{
		ID:      ID,
		Content: p.Args["content"].(string),
	}

	return dbQueries.UpdateComment(p.Context, params)
}

func DeleteComment(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)
	id := p.Args["id"].(string)

	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return dbQueries.DeleteComment(p.Context, ID)
}
