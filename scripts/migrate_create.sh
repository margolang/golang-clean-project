#!/usr/bin/env bash

goose -dir ./migrations create $1 sql
