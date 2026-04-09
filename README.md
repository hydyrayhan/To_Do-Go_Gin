🚀 Secure Go To-Do API
A high-performance, stateless RESTful API built with Golang and the Gin Gonic framework. This project implements a full task management system protected by JWT authentication, ensuring secure data access and efficient performance.

🔐 Key Features
JWT Authentication: Secure user registration and login flow with signed token issuance and validation.

RESTful CRUD: Full lifecycle management for tasks (Create, Read, Update, Delete).

Middleware Integration: Custom Gin middleware for handling authorized routes and token extraction.

Stateless Architecture: Designed for scalability using standard JWT patterns.

Performance-First: Leverages Go’s compiled speed and Gin’s lightweight routing engine.

🛠 Tech Stack
Language: Go

Web Framework: Gin Gonic

Security: JWT (JSON Web Tokens)

Data Handling: JSON serialization/deserialization

Some needed terminal codes
1. go init todo_api => initialize our project
2. go get -u github.com/gin-gonic/gin => Install gin framework

For run

1. go run cmd/api/main.go
2. go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
3. go get github.com/gin-gonic/gin@v1.10.0
4. go mod tidy


Migration commands
migrate create -ext sql -dir migrations -seq create_todos_table

./scripts/migrate.sh up        # this is create all table and if you specific table need then put number (order number of created table 000001) after up like up 1
./scripts/migrate.sh down      # this will delete last first by one by first order number 2 then 1
./scripts/migrate.sh create name_here  # create new migration

If dirty database error
./scripts/migrate.sh down
./scripts/migrate.sh up
