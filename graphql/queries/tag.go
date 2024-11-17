package queries

import (
	"github.com/AbdelilahOu/GoferQl/graphql/resolvers"
	"github.com/AbdelilahOu/GoferQl/graphql/types"

	"github.com/graphql-go/graphql"
)

var TagQueries = graphql.Fields{
	"tag": &graphql.Field{
		Type: types.TagType,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: resolvers.GetTag,
	},
	"tags": &graphql.Field{
		Type: graphql.NewList(types.TagType),
		Args: graphql.FieldConfigArgument{
			"limit":  &graphql.ArgumentConfig{Type: graphql.Int},
			"offset": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: resolvers.ListTags,
	},
}
