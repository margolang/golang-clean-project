#!/usr/bin/env bash
db='postgres://postgres:postgres@localhost:5432/presentation?sslmode=disable'

if [ -z "$1" ] 
then
    goose -dir ./migrations -allow-missing postgres $db up
    sqlboiler -c ./configs/sqlboiler.toml psql
else
    migrate -path ./migrations -database "$db" up "$1"
fi
