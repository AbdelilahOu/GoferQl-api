package resolvers

import (
	"fmt"

	db "github.com/AbdelilahOu/GoferQl/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

func ListTags(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	var Limit int32 = 20
	var Offset int32 = 0

	if val, ok := p.Args["limit"]; ok && val != nil {
		Limit = val.(int32)
	}
	if val, ok := p.Args["offset"]; ok && val != nil {
		Offset = val.(int32)
	}

	return dbQueries.ListTags(p.Context, db.ListTagsParams{
		Limit:  Limit,
		Offset: Offset,
	})
}

func GetTag(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	if val, ok := p.Args["id"]; ok && val != nil {
		ID, err := uuid.Parse(val.(string))
		if err != nil {
			return nil, err
		}
		return dbQueries.GetTag(p.Context, ID)
	}

	return nil, fmt.Errorf("id not provided")
}

func CreateTag(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	return dbQueries.CreateTag(p.Context, p.Args["name"].(string))
}
