# Chirpy

Chirpy is a sample HTTP server built with Go. It supports user authentication, chirp posting, and includes
administrative endpoints for metrics and system resets.


---

## Table of Contents

- [Features](#features)
- [Requirements](#requirements)
- [Installation](#installation)
- [Environment Variables](#environment-variables)
- [Database Migrations](#database-migrations)
- [Running the Server](#running-the-server)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [License](#license)

---

## Features

- User creation, login, and update using JWT authentication.
- Post chirps (short messages) with creation, retrieval, and deletion endpoints.
- Refreshing and revoking JWT tokens.
- Admin endpoints to check the server metrics and reset the hit counter.
- Webhook support for external events.

---

## Requirements

- Go 1.23
- PostgreSQL
- [Goose](https://github.com/pressly/goose) for handling database migrations
- [sqlc](https://github.com/sqlc-dev/sqlc) for handling generate type-safe code from SQL

---

## Installation

1. **Clone the repository:**

   ```bash
   git clone https://github.com/<username>/<repository>.git
   cd <repository>
   ```

2. **Install dependencies:**

   ```bash
   go mod tidy
   ```

3. **Setup the PostgreSQL database** and create a database named `chirpy`.

4. **Set up environment variables:**

   Copy the example below into a `.env` file at the root of the project.

---

## Environment Variables

Create a `.env` file at the root with the following configuration (replace placeholders with your actual values):

```bash
PLATFORM=dev
DB_URL=postgres://<USER>:<PASSWORD>@localhost:5432/chirpy
JWT_SECRET=<your_jwt_secret>
POLKA_KEY=<your_api_key>
```

---

## Database Migrations

Database migrations are managed with [goose](https://github.com/pressly/goose). To apply or revert migrations, follow
these steps:

1. **Navigate to the SQL schema directory:**

   ```bash
   cd sql/schema
   ```

2. **Migrate Up:**

   ```bash
   goose postgres "postgres://<USER>:<PASSWORD>@localhost:5432/chirpy" up
   ```

3. **Migrate Down:**

   ```bash
   goose postgres "postgres://<USER>:<PASSWORD>@localhost:5432/chirpy" down
   ```

4. **Write queries in SQL file :**
   ```bash
   cd sql/queries
   touch <name>.sql
   ```

5**generate type-safe code form SQL:**

   ```bash
   sqlc generate
   ```

---

## Running the Server

From the project root, run:

```bash
go run .
```

By default, the server listens on port **8080**. You can change this in the source if needed.

---

## API Endpoints

### User Endpoints

- **Create User:** `POST /api/users`
- **Update User:** `PUT /api/users`
- **Login:** `POST /api/login`
- **Refresh Token:** `POST /api/refresh`
- **Revoke Token:** `POST /api/revoke`

### Chirp Endpoints

- **Create Chirp:** `POST /api/chirps`
- **Retrieve Chirps:** `GET /api/chirps`
    - Optional query parameter `author_id` (UUID) to filter chirps by user.
    - Optional query parameter `sort` to specify the order in which results are returned. Use `asc` for ascending
      order (default) or `desc` for descending order.
- **Retrieve Single Chirp:** `GET /api/chirps/{chirpID}`
- **Delete Chirp:** `DELETE /api/chirps/{chirpID}`

### Admin Endpoints

- **Reset Hits:** `POST /admin/reset` (only available in dev mode)
- **Metrics:** `GET /admin/metrics`

### Webhook Endpoint

- **Polka Webhook:** `POST /api/polka/webhooks`

---

## Testing

Run the test suite with the following command:

```bash
go test ./...
```

---

## License

This project is licensed under the MIT License.
