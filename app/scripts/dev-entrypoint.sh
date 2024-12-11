#!/bin/sh

# Download dependencies
go mod download

# Build the application
go build -o ./cmd/web/main ./cmd/web/main.go

# Run migrations and seeders
go run ./cmd/web/main.go seed

# Run the application
go run ./cmd/web/main.go