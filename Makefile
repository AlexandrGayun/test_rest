#!/bin/env make

.PHONY: sqlc

sqlc:
	sqlc generate

#test:
#	GO_ENV=test SQL_DSN=postgres://postgres:secret@localhost:5432/entrello_test?sslmode=disable go test -count 1 -p 1 -cover -coverprofile coverage.out -coverpkg ./... ./...

test:
	go test -race -vet=off -v -count=1 ./...

down:
	docker compose down

up:
	docker compose up -d
