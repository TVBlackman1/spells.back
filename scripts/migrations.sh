#!/bin/bash

# used https://github.com/golang-migrate/migrate
# use sh from root of application


db_username=postgres
db_password=96pS1vO
host=tvblackman1.ru
port=5432
db_name=dnd_spells

migrate -database "postgres://$db_username:$db_password@$host:$port/$db_name?sslmode=disable" -path ./init/migrations up