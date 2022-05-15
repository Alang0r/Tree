#!/bin/sh

set -e

# echo "start db migration to $DB_SOURCE"
# /app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

# echo "start the app"

exec "$@"