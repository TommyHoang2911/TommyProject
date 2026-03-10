# Auth Service

This service provides user registration and authentication backed by PostgreSQL.

## Setup

1. Install PostgreSQL and create a database, e.g.: 

```sh
createdb authdb
```

2. Run the migration script:

```sh
psql -d authdb -f config/migrations.sql
```

3. Set `DATABASE_URL` environment variable (optional). Default is:
   `postgres://postgres:password@localhost:5432/authdb?sslmode=disable`.

4. Build and run the service:

```sh
go build ./...
./auth-service/server/main
```

The server listens on port `8080` by default.

## Registering a user

Send a POST request to `/register` with JSON body:

```json
{
  "email": "user@example.com",
  "password": "secret"
}
```

The user will be inserted into the database.

## Notes

- Passwords are stored in plaintext for now. Replace with hashing in production.
- Add more fields or validation as needed.
