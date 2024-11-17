package mutations

import (
	"github.com/AbdelilahOu/GoferQl/graphql/resolvers"
	"github.com/AbdelilahOu/GoferQl/graphql/types"

	"github.com/graphql-go/graphql"
)

var CommentMutations = graphql.Fields{
	"createComment": &graphql.Field{
		Type: types.CommentType,
		Args: graphql.FieldConfigArgument{
			"content":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"userId":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"postId":   &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"parentId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: resolvers.CreateComment,
	},
	"updateComment": &graphql.Field{
		Type: types.CommentType,
		Args: graphql.FieldConfigArgument{
			"id":      &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"content": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: resolvers.UpdateComment,
	},
	"deleteComment": &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: resolvers.DeleteComment,
	},
}
