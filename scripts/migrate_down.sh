#!/usr/bin/env bash
goose -dir ./migrations postgres "postgres://postgres:postgres@localhost:5432/presentation?sslmode=disable" down
sqlboiler -c ./configs/sqlboiler.toml psql
