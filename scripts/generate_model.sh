#!/usr/bin/env bash
db='postgres://umico_business:easytobruteforce@localhost:17654/umico_business?sslmode=disable'

docker run --rm -d -p 17654:5432 --name oms-migration-db -e POSTGRES_USER="umico_business" -e POSTGRES_PASSWORD="easytobruteforce" -e POSTGRES_DB="umico_business" postgres:11.11
sleep 5 # wait for db startup
goose -dir ./migrations postgres $db up
sed 's/= 7654/= 17654/g' configs/sqlboiler.toml > ./temp_boiler_config.toml
sqlboiler -c ./temp_boiler_config.toml psql
rm ./temp_boiler_config.toml
docker stop oms-migration-db