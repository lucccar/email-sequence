#!/bin/bash

# Start the test database container
docker compose -f docker-compose.test.yml up -d

# Wait for the database to be ready
sleep 7

# Run migrations
docker compose -f docker-compose.test.yml exec -T postgres-test sh -c "psql -U postgres -d mytestdatabase -f /migrations/001_create_sequences.up.sql"

# Run tests
go test ./... -v

# Stop and remove the test database container
# docker compose -f docker-compose.test.yml down
