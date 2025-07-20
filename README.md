# goserver

A simple Go server for managing users and chirps (short messages), with authentication, metrics, and webhooks. This project demonstrates a RESTful API with PostgreSQL integration and token-based authentication.

## Features
- User registration, login, update, and token refresh/revoke
- Chirp creation, retrieval, and deletion
- Metrics and readiness endpoints
- Webhook support
- Secure password storage and token management

## Prerequisites
- Go 1.20+
- PostgreSQL (default connection: `postgres://postgres:postgres@localhost:5432/chirpy`)
- [goose](https://github.com/pressly/goose) for database migrations

## Setup
1. **Clone the repository:**
   ```bash
   git clone <repo-url>
   cd goserver
   ```
2. **Install dependencies:**
   ```bash
   go mod tidy
   ```
3. **Configure your database:**
   Ensure PostgreSQL is running and accessible at the default connection string, or update the connection string in your environment.

## Database Migration
Run the following commands to apply database migrations:
```bash
pushd sql/schema
# Apply all migrations
goose postgres postgres://postgres:postgres@localhost:5432/chirpy up
popd
```

## Running the Server
Start the server with:
```bash
go run main.go
```

## API Overview
- `POST /api/users` - Register a new user
- `POST /api/login` - Login and receive tokens
- `PUT /api/users` - Update user info (auth required)
- `POST /api/refresh` - Refresh access token
- `POST /api/revoke` - Revoke refresh token
- `POST /api/chirps` - Create a chirp (auth required)
- `GET /api/chirps` - List chirps
- `DELETE /api/chirps/{id}` - Delete a chirp (auth required)
- `GET /admin/metrics` - Metrics endpoint
- `GET /admin/readiness` - Readiness probe
- `POST /api/webhooks` - Webhook endpoint

## Directory Structure
```
.
├── assets/                # Static assets (e.g., logo)
├── internal/
│   ├── auth/              # Authentication logic
│   └── database/          # Database models and queries
├── sql/
│   ├── queries/           # SQL query files
│   └── schema/            # Migration scripts
├── handler_*.go           # HTTP handlers
├── main.go                # Server entry point
├── user.go, json.go       # Utility and model files
├── README.md              # This file
```

## Testing
Run tests with:
```bash
go test ./...
```

## License
MIT License

---

# Original Quick Commands
## commands
```bash
pushd sql/schema
goose postgres postgres://postgres:postgres@localhost:5432/chirpy up
popd

pushd sql/schema
goose postgres postgres://postgres:postgres@localhost:5432/chirpy up
popd
```