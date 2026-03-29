# HRMS Backend (Go + Fiber)

A modular **HRMS (Human Resource Management System) backend** built using **Go**, **Fiber**, and **PostgreSQL**.  
This project follows a clean layered architecture so multiple developers can work on features independently.

---

# Tech Stack

- Go
- Fiber (HTTP Framework)
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
├── internal
│   └── onboarding            # Onboarding module
│       ├── handler
│       │   └── onboarding_handler.go
│       │
│       ├── service
│       │   └── onboarding_service.go
│       │
│       ├── repository
│       │   └── onboarding_repository.go
│       │
│       ├── model
│       │   └── employee.go
│       │
│       └── routes
│           └── onboarding_routes.go
│
├── pkg
│   ├── database              # Database connection & migrations
│   ├── middleware            # Fiber middleware
│   └── utils                 # Helper utilities
│
├── migrations                # SQL migrations
├── .env.example
├── go.mod
├── go.sum
└── README.md
```

---

# Architecture

Each module follows the **Handler → Service → Repository architecture**.

```
Client Request
      ↓
Fiber Router
      ↓
Handler
      ↓
Service
      ↓
Repository
      ↓
PostgreSQL
```

---

# Handler Layer

The **handler layer** manages HTTP requests and responses.

Responsibilities:

- Parse request body
- Read URL params
- Call service layer
- Send response to client

Example:

```go
package handler

import "github.com/gofiber/fiber/v2"

func Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Onboarding service working",
	})
}
```

---

# Service Layer

The **service layer** contains business logic.

Responsibilities:

- Implement business rules
- Validate data
- Call repository functions

Example:

```go
package service

import "hrms/internal/onboarding/repository"

type OnboardingService struct {
	Repo *repository.OnboardingRepository
}

func (s *OnboardingService) CreateEmployee(name string) error {
	return s.Repo.InsertEmployee(name)
}
```

---

# Repository Layer

The **repository layer** communicates with the database.

Responsibilities:

- Execute SQL queries
- Fetch data
- Insert / update / delete records

Example:

```go
package repository

import "database/sql"

type OnboardingRepository struct {
	DB *sql.DB
}

func (r *OnboardingRepository) InsertEmployee(name string) error {
	query := `INSERT INTO employees(name) VALUES($1)`
	_, err := r.DB.Exec(query, name)
	return err
}
```

---

# Model Layer

Models represent database entities.

Example:

```go
package model

type Employee struct {
	ID         int
	FirstName  string
	LastName   string
	Email      string
	Department string
}
```

---

# Routes Layer

Routes connect endpoints to handlers.

Example:

```go
package routes

import (
	"github.com/gofiber/fiber/v2"
	"hrms/internal/onboarding/handler"
)

func RegisterOnboardingRoutes(app *fiber.App) {

	onboarding := app.Group("/api/onboarding")

	h := handler.OnboardingHandler{}

	onboarding.Get("/health", h.Health)
}
```

Example endpoint:

```
GET /api/onboarding/health
```

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

Create `.env` from example.

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

```
go mod tidy
```

---

# Setup PostgreSQL Database

Create database:

```
CREATE DATABASE hrms;
```

Verify:

```
\l
```

---

# Run the Server

From project root:

```
go run cmd/server/main.go
```

Server will start at:

```
http://localhost:8080
```

---

# Test API

Health check endpoint:

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

Create a new feature branch:

```
git checkout -b feature/<feature-name>
```

Example:

```
git checkout -b feature/onboarding-api
```

Commit changes:

```
git add .
git commit -m "Added onboarding API"
```

Push branch:

```
git push origin feature/onboarding-api
```

Then open a **Pull Request**.

---

# Coding Guidelines

- Keep handlers thin
- Business logic belongs in services
- SQL queries belong in repositories
- Follow Go formatting rules

Format code:

```
gofmt -w .
```

Run linter:

```
golangci-lint run
```

---

# Quick Project Setup

```
git clone <repo-url>

cd hrms

cp .env.example .env

go mod tidy

go run cmd/server/main.go
```

---