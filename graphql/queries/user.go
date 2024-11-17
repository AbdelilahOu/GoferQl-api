package queries

import (
	"github.com/AbdelilahOu/GoferQl/graphql/resolvers"
	"github.com/AbdelilahOu/GoferQl/graphql/types"

	"github.com/graphql-go/graphql"
)

var UserQueries = graphql.Fields{
	"user": &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"id":    &graphql.ArgumentConfig{Type: graphql.String},
			"email": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: resolvers.GetUser,
	},
	"users": &graphql.Field{
		Type: graphql.NewList(types.UserType),
		Args: graphql.FieldConfigArgument{
			"limit":  &graphql.ArgumentConfig{Type: graphql.Int},
			"offset": &graphql.ArgumentConfig{Type: graphql.Int},
		},
		Resolve: resolvers.ListUsers,
	},
}
