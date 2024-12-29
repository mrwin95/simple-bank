#!/bin/sh

set -e

echo "Start db migration"

/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "Start app"

exec "$@"
