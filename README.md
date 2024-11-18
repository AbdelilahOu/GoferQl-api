# GoferQL API

A GraphQL server implementation built with Go, PostgreSQL, and github.com/graphql-go/graphql.

## Overview

GoferQL API is a GraphQL server that provides a flexible API for managing posts, comments, users, categories, and tags. It uses SQLC for type-safe database operations and the graphql-go package for GraphQL implementation.

## Tech Stack

- Go
- PostgreSQL
- github.com/graphql-go/graphql
- SQLC

## Project Structure

```
.
├── cmd/
│   ├── seed/
│   │   └── main.go // generate dummy data
│   └── server/
│       └── main.go // server handler
├── config/
│   └── config.go // server config
├── graphql/
│   ├── mutations/  // mutations
│   │   ├── category.go
│   │   ├── comment.go
│   │   ├── post.go
│   │   ├── postTag.go
│   │   ├── tag.go
│   │   └── user.go
│   ├── queries/  //queries
│   │   ├── category.go
│   │   ├── comment.go
│   │   ├── post.go
│   │   ├── tag.go
│   │   └── user.go
│   ├── resolvers/  // resolvers
│   │   ├── category.go
│   │   ├── comment.go
│   │   ├── post.go
│   │   ├── postTag.go
│   │   ├── tag.go
│   │   └── user.go
│   ├── types/  // types
│   │   └── types.go
│   └── utils/  // schema
│       └── schema.go
├── internal/
│   └── db/ // database + sqlc
│       ├── migrations/
│       ├── queries/
│       └── sqlc/
├── .env
├── go.mod
├── go.sum
├── Makefile // all project commands
├── README.md
├── schema.graphql // graphql schema
└── sqlc.yaml // sqlc config
```

## Setup

1. Clone the repository:

```bash
git clone https://github.com/AbdelilahOu/GoferQl-api
cd GoferQl-api
```

2. Install dependencies:

```bash
go mod download
```

3. Set up the database:

```bash
# Create database container
make containerup

# Create PostgreSQL database
make createdb

# Run migrations
make migrations-up
```

4. seed database:

```bash
make seed
```

4. Configure environment variables:

```bash
cp .env.example .env
# Edit .env with your configuration
```

5. Run the server:

```bash
make server
```

## GraphQL Schema

The API supports the following main types:

- User (username, email, bio)
- Post (title, content, status)
- Comment (content, nested replies)
- Category (name, description)
- Tag (name)

### Example Queries

Fetch posts with author and comments:

```graphql
query {
  posts(limit: 10, offset: 0) {
    id
    title
    content
    user {
      username
    }
    comments {
      content
      user {
        username
      }
    }
  }
}
```

Create a new post:

```graphql
mutation {
  createPost(
    title: "Hello World"
    content: "This is my first post"
    status: "published"
    userId: "user-id"
    categoryId: "category-id"
  ) {
    id
    title
    createdAt
  }
}
```
