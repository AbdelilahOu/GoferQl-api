package mutations

import (
	"github.com/AbdelilahOu/GoferQl/graphql/resolvers"
	"github.com/AbdelilahOu/GoferQl/graphql/types"

	"github.com/graphql-go/graphql"
)

var UserMutations = graphql.Fields{
	"createUser": &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"username": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"email":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"password": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"bio":      &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: resolvers.CreateUser,
	},
	"updateUser": &graphql.Field{
		Type: types.UserType,
		Args: graphql.FieldConfigArgument{
			"id":       &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"username": &graphql.ArgumentConfig{Type: graphql.String},
			"email":    &graphql.ArgumentConfig{Type: graphql.String},
			"bio":      &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: resolvers.UpdateUser,
	},
	"deleteUser": &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: resolvers.DeleteUser,
	},
}
