package queries

import (
	"github.com/AbdelilahOu/GoferQl/graphql/resolvers"
	"github.com/AbdelilahOu/GoferQl/graphql/types"

	"github.com/graphql-go/graphql"
)

var CommentQueries = graphql.Fields{
	"comments": &graphql.Field{
		Type: graphql.NewList(types.CommentType),
		Args: graphql.FieldConfigArgument{
			"postId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: resolvers.ListCommentsByPostID,
	},
}
