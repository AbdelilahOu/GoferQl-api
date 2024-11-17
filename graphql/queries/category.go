package queries

import (
	"github.com/AbdelilahOu/GoferQl/graphql/resolvers"
	"github.com/AbdelilahOu/GoferQl/graphql/types"

	"github.com/graphql-go/graphql"
)

var CategoryQueries = graphql.Fields{
	"category": &graphql.Field{
		Type: types.CategoryType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: resolvers.GetCategory,
	},
	"categories": &graphql.Field{
		Type: graphql.NewList(types.CategoryType),
		Args: graphql.FieldConfigArgument{
			"limit":  &graphql.ArgumentConfig{Type: graphql.Int},
			"offset": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: resolvers.ListCategories,
	},
}
