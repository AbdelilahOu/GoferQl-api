package queries

import (
	"github.com/AbdelilahOu/GoferQl/graphql/resolvers"
	"github.com/AbdelilahOu/GoferQl/graphql/types"

	"github.com/graphql-go/graphql"
)

var PostTagsQueries = graphql.Fields{
	"postTags": &graphql.Field{
		Type: graphql.NewList(types.PostTagType),
		Args: graphql.FieldConfigArgument{
			"postId": &graphql.ArgumentConfig{Type: graphql.String},
			"tagId":  &graphql.ArgumentConfig{Type: graphql.String},
		},
		Resolve: resolvers.ListPostTags,
	},
}
