# Numinaut Backend

A Go-based REST API service for Numinaut

## Features

- RESTful API endpoints for expressions
- Filtering expressions by multiple criteria
- PostgreSQL database for persistent storage
- Structured JSON logging with Zap
- Hot reloading for development

## Tech Stack

- **Go** (v1.23.0)
- **PostgreSQL** - Database
- **pgx** - PostgreSQL driver
- **sqlc** - Type-safe SQL in Go
- **gorilla/mux** - HTTP router
- **zap** - Structured logging
- **godotenv** - Environment configuration
- **Air** - Live reloading for development
- **goose** - Database migration tool

## Project Structure

numinaut-be/
├── cmd/api/ # Application entry point
├── internal/ # Internal application code
│ ├── handler/ # HTTP request handlers
│ ├── middleware/ # HTTP middleware components
│ ├── repository/ # Database access layer
│ └── service/ # Business logic layer
├── migrations/ # Database migration files
├── pkg/ # Shared packages
│ └── logger/ # Logging utilities
├── query/ # SQL queries for sqlc
├── .air.toml # Air configuration for hot reloading
├── .env # Environment variables
└── sqlc.yaml # sqlc configuration

## Getting Started

### Prerequisites

- Go 1.23.0 or higher
- PostgreSQL
- Air (optional, for hot reloading)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/vishwaszadte/numinaut-be.git
   cd numinaut-be

   ```

2. Install dependencies:

```bash
go mod download
```

3. Set up your environment variables:

```plaintext
POSTGRESQL_URL=postgres://postgres:pass1234@localhost:5432/numinaut?sslmode=disable
GOOSE_DRIVER=postgres
GOOSE_DBSTRING=postgres://postgres:pass1234@localhost:5432/numinaut?sslmode=disable
GOOSE_MIGRATION_DIR=migrations
LOG_LEVEL=debug
```

4. Run migrations:

```bash
goose up
```

### Running the Application

- Standard Run

```bash
go run cmd/api/main.go
```

- Development with Hot Reload

```bash
air
```

The server will start on `localhost:8080`.
