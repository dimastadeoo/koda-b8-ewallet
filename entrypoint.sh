#!/bin/sh

set -e

# Jalankan PostgreSQL di background
docker-entrypoint.sh postgres &

# Tunggu PostgreSQL siap
until pg_isready \
    -h localhost \
    -U "$POSTGRES_USER" \
    -d "$POSTGRES_DB"
do
    sleep 1
done

# Buat file .env untuk aplikasi Go
cat <<EOF >/app/.env
DATABASE_URL=postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@localhost:5432/${POSTGRES_DB}?sslmode=disable
EOF

echo "Database is ready."
echo "Starting E-Wallet..."

exec /app/program