#!/bin/sh

# ensures this script will exit immediately if a command returns a non-zero status
set -e

# echo "run db migration"
# source /app/app.env
# /app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
# this means take all paramets passed to this script and run it
exec "$@"