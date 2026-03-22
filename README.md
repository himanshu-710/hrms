# HRMS Backend (Go + Gin)

A modular **HRMS (Human Resource Management System) backend** built using **Go**, **Gin**, and **PostgreSQL**.  
This project is structured so that multiple teams can work on different HRMS modules in parallel.

---

# Tech Stack

- Go
- Gin (HTTP Framework)
- PostgreSQL
- Zap Logger
- Godotenv

---

# Project Structure

```
hrms
│
├── cmd
│   └── server
│       └── main.go           # Application entry point
│
├── config                    # Environment configuration
│
├── internal                  # Feature modules
│   ├── attendance
│   ├── leave
│   ├── onboarding
│   ├── expenses
│   └── assets
│
├── pkg                       # Shared utilities
│   ├── database
│   ├── middleware
│   └── utils
│
├── .env.example              # Example environment variables
├── go.mod
├── go.sum
└── README.md
```

Each feature module follows this architecture:

```
module
│
├── handler       # HTTP request handlers
├── service       # Business logic
└── repository    # Database queries
```

Architecture Flow:

```
Handler → Service → Repository → Database
```
---

# Module Implementation Guide

Each module inside `internal/` should follow the **Handler → Service → Repository architecture**.

Example module:

```
internal/onboarding
│
├── handler
│   └── onboarding_handler.go
│
├── service
│   └── onboarding_service.go
│
├── repository
│   └── onboarding_repository.go
│
├── model
│   └── onboarding_model.go
│
└── routes
    └── onboarding_routes.go
```

---

# Handler Layer (HTTP Layer)

The **handler** is responsible for:

- Receiving HTTP requests
- Parsing request body / params
- Calling the service layer
- Sending the HTTP response

Handlers should **not contain business logic or database queries**.

Example:

```go
package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"hrms/internal/onboarding/service"
)

type OnboardingHandler struct {
	Service service.OnboardingService
}

func (h *OnboardingHandler) Health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Onboarding service working",
	})
}
```

Responsibilities:

- Handle HTTP requests
- Validate input
- Call service functions
- Return response

---

# Service Layer (Business Logic)

The **service layer** contains the business logic of the application.

Responsibilities:

- Implement business rules
- Coordinate between handler and repository
- Validate business logic

Example:

```go
package service

import "hrms/internal/onboarding/repository"

type OnboardingService struct {
	Repo repository.OnboardingRepository
}

func (s *OnboardingService) CreateEmployee(name string) error {
	return s.Repo.InsertEmployee(name)
}
```

Responsibilities:

- Business logic
- Data processing
- Calling repository functions

---

# Repository Layer (Database Layer)

The **repository layer** interacts directly with the database.

Responsibilities:

- SQL queries
- Data persistence
- Returning data from DB

Example:

```go
package repository

import (
	"database/sql"
)

type OnboardingRepository struct {
	DB *sql.DB
}

func (r *OnboardingRepository) InsertEmployee(name string) error {
	query := `INSERT INTO employees(name) VALUES($1)`
	_, err := r.DB.Exec(query, name)
	return err
}
```

Responsibilities:

- Execute SQL queries
- Fetch data
- Insert/update/delete records

---

# Model Layer (Database Models)

Models define the **structure of database entities**.

Example:

```go
package model

type Employee struct {
	ID        int
	Name      string
	Email     string
	Phone     string
	Department string
}
```

Responsibilities:

- Define entity structures
- Represent database tables
- Used for request/response mapping

---

# Routes Layer

Routes connect HTTP endpoints to handlers.

Example:

```go
package routes

import (
	"github.com/gin-gonic/gin"
	"hrms/internal/onboarding/handler"
)

func RegisterOnboardingRoutes(r *gin.RouterGroup) {

	onboarding := r.Group("/onboarding")

	h := handler.OnboardingHandler{}

	onboarding.GET("/health", h.Health)
}
```

Responsibilities:

- Register endpoints
- Map endpoints to handlers

Example endpoint:

```
GET /api/onboarding/health
```

---

# Request Flow

The request flow inside the application:

```
Client Request
      ↓
Gin Router
      ↓
Handler
      ↓
Service
      ↓
Repository
      ↓
PostgreSQL Database
```

---

# Best Practices

- Keep handlers **thin**
- Business logic should be in **services**
- Database queries should only be in **repositories**
- Do not access database directly from handlers
- Use models for request and response structures

---

# Example API Implementation Flow

Example endpoint:

```
POST /api/onboarding/employee
```

Flow:

```
Handler → Service → Repository → Database
```

1. Handler receives HTTP request
2. Service processes business logic
3. Repository executes database query
4. Response returned to client
---

# Clone the Repository

```
git clone <repo-url>
cd hrms
```

Example:

```
git clone https://github.com/<username>/hrms.git
cd hrms
```

---

# Setup Environment Variables

Create a `.env` file from the example.

```
cp .env.example .env
```

Example `.env`:

```
PORT=8080

DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=hrms
DB_SSLMODE=disable
```

---

# Install Dependencies

Run:

```
go mod tidy
```

This installs all required Go packages.

---

# Setup PostgreSQL Database

Create the database:

```
CREATE DATABASE hrms;
```

Verify it exists:

```
\l
```

---

# Run the Server

From the project root:

```
go run cmd/server/main.go
```

Server will start at:

```
http://localhost:8080
```

---

# Test API

Example test endpoint:

```
GET http://localhost:8080/api/onboarding/health
```

Expected response:

```json
{
  "message": "Onboarding service working"
}
```


---

# Development Workflow

Create a feature branch before working.

```
git checkout -b feature/<feature-name>
```

Example:

```
git checkout -b feature/onboarding-api
```

Commit your changes:

```
git add .
git commit -m "Added onboarding API"
```

Push your branch:

```
git push origin feature/onboarding-api
```

Then open a **Pull Request**.

---

# Coding Guidelines

- Follow Go formatting rules

```
gofmt -w .
```

- Run lint checks

```
golangci-lint run
```

- Keep handlers small
- Business logic should be in **service layer**
- Database queries should be in **repository layer**

---

# Running the Project (Quick Setup)

```
git clone <repo-url>

cd hrms

cp .env.example .env

go mod tidy

go run cmd/server/main.go
```

---

