#!make
include .env
CURRENT_DIR := $(abspath $(dir $(lastword $(MAKEFILE_LIST))))

containerup:
	docker run --name go-graphql-db -e POSTGRES_USER=${DB_USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -p 5432:5432 -d postgres:15

createdb: 
	docker exec -it go-graphql-db createdb --username=${DB_USER} --owner=${DB_USER} ${DB_NAME}

dropdb:
	docker exec -it go-graphql-db dropdb ${DB_NAME}

migrations-up:
	cd internal/db/migrations && goose postgres ${DB_URL} up

migration-down:
	cd internal/db/migrations && goose postgres ${DB_URL} down

migration-new:
	cd internal/db/migrations && goose postgres ${DB_URL} create $(name) sql

sqlc: 
	docker run --rm -v ${CURRENT_DIR}:/src -w /src sqlc/sqlc generate

server: 
	go run cmd/server/main.go
