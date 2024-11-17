package types

import "github.com/graphql-go/graphql"

var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.String},
			"username":  &graphql.Field{Type: graphql.String},
			"email":     &graphql.Field{Type: graphql.String},
			"bio":       &graphql.Field{Type: graphql.String},
			"createdAt": &graphql.Field{Type: graphql.DateTime},
			"password":  &graphql.Field{Type: graphql.String},
		},
	},
)

var CategoryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Category",
		Fields: graphql.Fields{
			"id":          &graphql.Field{Type: graphql.String},
			"name":        &graphql.Field{Type: graphql.String},
			"description": &graphql.Field{Type: graphql.String},
			"createdAt":   &graphql.Field{Type: graphql.DateTime},
		},
	},
)

var PostType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Post",
		Fields: graphql.Fields{
			"id":         &graphql.Field{Type: graphql.String},
			"title":      &graphql.Field{Type: graphql.String},
			"content":    &graphql.Field{Type: graphql.String},
			"userId":     &graphql.Field{Type: graphql.String},
			"categoryId": &graphql.Field{Type: graphql.String},
			"status":     &graphql.Field{Type: graphql.String},
			"createdAt":  &graphql.Field{Type: graphql.DateTime},
			"updatedAt":  &graphql.Field{Type: graphql.DateTime},
		},
	},
)

var CommentType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Comment",
		Fields: graphql.Fields{
			"id":        &graphql.Field{Type: graphql.String},
			"content":   &graphql.Field{Type: graphql.String},
			"userId":    &graphql.Field{Type: graphql.String},
			"postId":    &graphql.Field{Type: graphql.String},
			"parentId":  &graphql.Field{Type: graphql.String},
			"createdAt": &graphql.Field{Type: graphql.DateTime},
		},
	},
)

var TagType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Tag",
		Fields: graphql.Fields{
			"id":   &graphql.Field{Type: graphql.String},
			"name": &graphql.Field{Type: graphql.String},
		},
	},
)

var PostTagType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "PostTag",
		Fields: graphql.Fields{
			"postId": &graphql.Field{Type: graphql.String},
			"tagId":  &graphql.Field{Type: graphql.String},
		},
	},
)
