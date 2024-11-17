// graph/schema.go
package graph

import (
	"github.com/AbdelilahOu/GoferQl/graphql/mutations"
	"github.com/AbdelilahOu/GoferQl/graphql/queries"
	"github.com/graphql-go/graphql"
)

func NewSchema() (graphql.Schema, error) {
	QueryFields := graphql.Fields{}
	for key, val := range queries.UserQueries {
		QueryFields[key] = val
	}
	for key, val := range queries.CategoryQueries {
		QueryFields[key] = val
	}

	MutationFields := graphql.Fields{}
	for key, val := range mutations.UserMutations {
		MutationFields[key] = val
	}
	for key, val := range mutations.CategoryMutations {
		MutationFields[key] = val
	}

	return graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   "Query",
					Fields: QueryFields,
				},
			),
			Mutation: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   "Mutation",
					Fields: MutationFields,
				},
			),
		},
	)
}
