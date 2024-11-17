package resolvers

import (
	"fmt"

	"github.com/AbdelilahOu/GoferQl/graphql/utils"
	db "github.com/AbdelilahOu/GoferQl/internal/db/sqlc"
	"github.com/google/uuid"
	"github.com/graphql-go/graphql"
)

func ListCategories(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	var Limit int = 20
	var Offset int = 0

	if val, ok := p.Args["limit"]; ok && val != nil {
		Limit = val.(int)
	}
	if val, ok := p.Args["offset"]; ok && val != nil {
		Offset = val.(int)
	}

	return dbQueries.ListCategories(p.Context, db.ListCategoriesParams{
		Limit:  int32(Limit),
		Offset: int32(Offset),
	})
}

func GetCategory(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	if val, ok := p.Args["id"]; ok && val != nil {
		ID, err := uuid.Parse(val.(string))
		if err != nil {
			return nil, err
		}
		return dbQueries.GetCategory(p.Context, ID)
	}

	return nil, fmt.Errorf("id not provided")
}

func CreateCategory(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	params := db.CreateCategoryParams{
		Name:        p.Args["name"].(string),
		Description: utils.NullablePgTypeText(p.Args, "description"),
	}

	return dbQueries.CreateCategory(p.Context, params)
}

func UpdateCategory(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)

	ID, err := uuid.Parse(p.Args["id"].(string))
	if err != nil {
		return nil, err
	}

	params := db.UpdateCategoryParams{
		ID:          ID,
		Name:        utils.NullablePgTypeText(p.Args, "name"),
		Description: utils.NullablePgTypeText(p.Args, "description"),
	}

	return dbQueries.UpdateCategory(p.Context, params)
}

func DeleteCategory(p graphql.ResolveParams) (interface{}, error) {
	dbQueries := p.Context.Value("db").(*db.Queries)
	id := p.Args["id"].(string)

	ID, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}

	return dbQueries.DeleteCategory(p.Context, ID)
}
