package mutations

import (
	"github.com/AbdelilahOu/GoferQl/graphql/resolvers"
	"github.com/AbdelilahOu/GoferQl/graphql/types"

	"github.com/graphql-go/graphql"
)

var TagMutations = graphql.Fields{
	"createTag": &graphql.Field{
		Type: types.TagType,
		Args: graphql.FieldConfigArgument{
			"name": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
		},
		Resolve: resolvers.CreateTag,
	},
}
