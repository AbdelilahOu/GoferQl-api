package mutations

import (
	"github.com/AbdelilahOu/GoferQl/graphql/resolvers"
	"github.com/AbdelilahOu/GoferQl/graphql/types"

	"github.com/graphql-go/graphql"
)

var PostMutations = graphql.Fields{
	"createPost": &graphql.Field{
		Type: types.PostType,
		Args: graphql.FieldConfigArgument{
			"title":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"content":    &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"status":     &graphql.ArgumentConfig{Type: graphql.String},
			"userId":     &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"categoryId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: resolvers.CreatePost,
	},
	"updatePost": &graphql.Field{
		Type: types.PostType,
		Args: graphql.FieldConfigArgument{
			"id":         &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"title":      &graphql.ArgumentConfig{Type: graphql.String},
			"content":    &graphql.ArgumentConfig{Type: graphql.String},
			"status":     &graphql.ArgumentConfig{Type: graphql.String},
			"categoryId": &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: resolvers.UpdatePost,
	},
	"deletePost": &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: resolvers.DeletePost,
	},
}
