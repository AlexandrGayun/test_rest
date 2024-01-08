#!/bin/env make

.PHONY: sqlc

sqlc:
	sqlc generate

test:
	go test -race -vet=off -v -count=1 ./...

down:
	docker compose down

up:
	docker compose up
