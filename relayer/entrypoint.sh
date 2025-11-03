#!/bin/sh
set -e

echo "🔧 Starting database migration process..."

# Wait for PostgreSQL to be ready
until pg_isready -h postgres -p 5432 -U bridge_user -d bridge_local; do
  echo "⏳ Waiting for database to be ready..."
  sleep 2
done

echo "✅ Database is ready. Applying migrations..."

# Apply migrations using golang-migrate
migrate -path /app/migrations -database "postgres://bridge_user:bridge_password@postgres:5432/bridge_local?sslmode=disable" up

echo "🎉 Migrations applied successfully. Starting application..."

# Start the main application
exec ./relayer-bin --env staging
