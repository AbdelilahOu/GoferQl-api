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
	for key, val := range queries.PostQueries {
		QueryFields[key] = val
	}
	for key, val := range queries.CommentQueries {
		QueryFields[key] = val
	}
	for key, val := range queries.TagQueries {
		QueryFields[key] = val
	}

	MutationFields := graphql.Fields{}
	for key, val := range mutations.UserMutations {
		MutationFields[key] = val
	}
	for key, val := range mutations.CategoryMutations {
		MutationFields[key] = val
	}
	for key, val := range mutations.PostMutations {
		MutationFields[key] = val
	}
	for key, val := range mutations.CommentMutations {
		MutationFields[key] = val
	}
	for key, val := range mutations.TagMutations {
		MutationFields[key] = val
	}
	for key, val := range mutations.PostTagMutations {
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
