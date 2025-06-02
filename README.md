# User CRUD API with Gin and PostgreSQL

This is a simple REST API built with Gin framework and PostgreSQL database for managing users.

## Prerequisites

- Go 1.21 or higher
- PostgreSQL
- Git

## Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd user-crud-api
```

2. Install dependencies:
```bash
go mod download
```

3. Create PostgreSQL database:
```sql
CREATE DATABASE userdb;
```

4. Create a `.env` file in the root directory with the following content:
```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=userdb
DB_SSL_MODE=disable
SERVER_PORT=9191
```

5. Run the database migration:
```bash
psql -U postgres -d userdb -f migrations/001_init.sql
```

## Running the Application

```bash
go run main.go
```

The server will start on the port specified in your `.env` file (default: 9191)

## API Endpoints

### Create User
- **POST** `/users`
- Request Body:
```json
{
    "username": "john_doe",
    "email": "john@example.com",
    "password": "secretpassword"
}
```

### Get All Users
- **GET** `/users`

### Get User by ID
- **GET** `/users/:id`

### Update User
- **PUT** `/users/:id`
- Request Body:
```json
{
    "username": "john_doe_updated",
    "email": "john.updated@example.com",
    "password": "newpassword"
}
```

### Delete User
- **DELETE** `/users/:id`

## Error Handling

The API returns appropriate HTTP status codes and error messages in JSON format:

```json
{
    "error": "error message"
}
```

## Security Notes

- In a production environment, you should:
  - Use environment variables for database credentials (already implemented)
  - Implement proper password hashing
  - Add authentication and authorization
  - Use HTTPS
  - Implement rate limiting
  - Add input validation
  - Keep your `.env` file secure and never commit it to version control
  - Use strong passwords for your database
  - Enable SSL mode for database connections in production