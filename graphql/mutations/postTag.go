package mutations

import (
	"github.com/AbdelilahOu/GoferQl/graphql/resolvers"
	"github.com/AbdelilahOu/GoferQl/graphql/types"

	"github.com/graphql-go/graphql"
)

var PostTagMutations = graphql.Fields{
	"createPostTag": &graphql.Field{
		Type: types.PostTagType,
		Args: graphql.FieldConfigArgument{
			"postId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"tagId":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: resolvers.AddPostTag,
	},
	"removePostTag": &graphql.Field{
		Type: graphql.Boolean,
		Args: graphql.FieldConfigArgument{
			"postId": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			"tagId":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: resolvers.RemovePostTag,
	},
}
