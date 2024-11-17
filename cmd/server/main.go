package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/AbdelilahOu/GoferQl/config"
	graph "github.com/AbdelilahOu/GoferQl/graphql"
	db "github.com/AbdelilahOu/GoferQl/internal/db/sqlc"
	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/graphql-go/graphql"
)

func main() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	pgPool, err := pgxpool.New(context.Background(), config.DBUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	defer pgPool.Close()

	if err = pgPool.Ping(context.Background()); err != nil {
		log.Fatal("coudnt ping db:", err)
	}

	schema, err := graph.NewSchema()
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := db.New(pgPool)

	http.HandleFunc("/graphql", createGraphQLHandler(schema, dbQueries))
	if err := http.ListenAndServe(fmt.Sprintf(":%s", config.PORT), nil); err != nil {
		log.Fatal("server error:", err)
	}
}

func createGraphQLHandler(schema graphql.Schema, dbQueries *db.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var params struct {
			Query         string                 `json:"query"`
			OperationName string                 `json:"operationName"`
			Variables     map[string]interface{} `json:"variables"`
		}
		if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		result := graphql.Do(graphql.Params{
			Schema:        schema,
			RequestString: params.Query,
			OperationName: params.OperationName,
			Context:       context.WithValue(r.Context(), "db", dbQueries),
		})

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
