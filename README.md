# 🚀 Secure Go To-Do API

[![Go Version](https://img.shields.io/badge/Go-1.25.0-blue.svg)](https://golang.org/)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

A high-performance, stateless RESTful API built with Golang and the Gin Gonic framework. This project implements a full task management system protected by JWT authentication, ensuring secure data access and efficient performance.

## 🔐 Key Features

- **JWT Authentication**: Secure user registration and login flow with signed token issuance and validation
- **RESTful CRUD Operations**: Full lifecycle management for tasks (Create, Read, Update, Delete)
- **Middleware Integration**: Custom Gin middleware for handling authorized routes and token extraction
- **Stateless Architecture**: Designed for scalability using standard JWT patterns
- **Performance-First**: Leverages Go's compiled speed and Gin's lightweight routing engine
- **PostgreSQL Database**: Robust data persistence with user-specific todo isolation
- **Migration Support**: Database schema management with golang-migrate

## 🛠 Tech Stack

- **Language**: Go 1.25.0
- **Web Framework**: Gin Gonic
- **Security**: JWT (JSON Web Tokens)
- **Database**: PostgreSQL with pgx driver
- **Configuration**: Environment variables with godotenv
- **Migration Tool**: golang-migrate

## 📋 Prerequisites

- Go 1.25.0 or later
- PostgreSQL database
- golang-migrate CLI tool (for database migrations)

## 🚀 Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/todo_api.git
   cd todo_api
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   ```
   Edit `.env` with your configuration:
   ```env
   DATABASE_URL=postgres://username:password@localhost:5432/todo_db?sslmode=disable
   JWT_SECRET=your_super_secret_jwt_key_here
   PORT=8080
   ```

4. **Run database migrations**
   ```bash
   ./scripts/migrate.sh up
   ```

## 🏃‍♂️ Running the Application

### Development Mode
```bash
go run cmd/api/main.go
```

### Build and Run
```bash
go build -o bin/todo_api cmd/api/main.go
./bin/todo_api
```

The server will start on the port specified in your `.env` file (default: 8080).

## 📚 API Documentation

### Authentication Endpoints

#### Register User
```http
POST /auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

#### Login User
```http
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

**Response:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

### Todo Endpoints

All todo endpoints require authentication. Include the JWT token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

#### Create Todo
```http
POST /todos
Content-Type: application/json

{
  "title": "Buy groceries",
  "completed": false
}
```

#### Get All Todos
```http
GET /todos
```

#### Get Todo by ID
```http
GET /todos/:id
```

#### Update Todo
```http
PUT /todos/:id
Content-Type: application/json

{
  "title": "Buy groceries and cook dinner",
  "completed": true
}
```

#### Delete Todo
```http
DELETE /todos/:id
```

## 🗄️ Database Management

### Migration Commands

Create a new migration:
```bash
migrate create -ext sql -dir migrations -seq create_table_name
```

Run migrations:
```bash
./scripts/migrate.sh up        # Apply all pending migrations
./scripts/migrate.sh up 1      # Apply specific migration (e.g., 000001)
```

Rollback migrations:
```bash
./scripts/migrate.sh down      # Rollback last migration
./scripts/migrate.sh down 1    # Rollback specific number of migrations
```

If you encounter a "dirty database" error:
```bash
./scripts/migrate.sh down
./scripts/migrate.sh up
```

## 🧪 Testing

```bash
go test ./...
```

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## 📝 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Gin Gonic](https://gin-gonic.com/) - The web framework used
- [golang-jwt](https://github.com/golang-jwt/jwt) - JWT implementation
- [pgx](https://github.com/jackc/pgx) - PostgreSQL driver
- [golang-migrate](https://github.com/golang-migrate/migrate) - Migration tool

## 📞 Support

If you have any questions or issues, please open an issue on GitHub or contact the maintainers.