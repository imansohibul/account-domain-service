## 🏗️ Project Structure & Software Architecture

This project follows a **modular clean architecture** pattern. It ensures high maintainability, testability, and clear separation of concerns.

# Account Domain Service

![Go Version](https://img.shields.io/badge/go-1.24%2B-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17.4-blue)
[![Docker](https://img.shields.io/badge/docker-24.0-blue?logo=docker)](https://www.docker.com/)
[![Docker Compose](https://img.shields.io/badge/docker--compose-3.9-blue)](https://docs.docker.com/compose/compose-file/compose-versioning/#version-39)


## 📂 Project Structure

```text
.
📦 account-domain-service
├── build/                   # Dockerfile for account service
│   └── Dockerfile           # Docker build context for account service
├── cmd/                     # Application entrypoints
│   └── main.go              # Main function as entrypoint for REST API, consumer, cron-job, etc
|   └── restapi.go           # Starts REST API
├── config/                  # Configuration management and dependency injection
├── db/
│   └── migrate/             # DB migrations using golang-migrate (up/down SQL files)
├── entity/                  # Domain entities and business rules
├── internal/
│   ├── repository/          # Data access layer (Postgres, etc.)
│   ├── rest/
│   │   ├── handler/         # Echo handlers (controllers)
|   |   ├── middleware/      # Custom middleware if any
│   │   ├── server/          # Server setup, routing, and middleware
│   └── usecase/             # Application use cases (interactors)
├── util/                    # Helper functions (e.g. validator, logging, formatting)
├── docker-compose.yaml      # Defines services (API, DB) for deployment
├── .env.sample              # Sample environment configuration
└── Makefile                 # Common scripts for building, testing, running
```

### 🧱 Architectural Layers

| Layer        | Responsibility |
|--------------|----------------|
| **Entity**   | Core domain logic: models and business rules. No framework or external dependency here. |
| **Usecase**  | Orchestrates application flow: how data moves and is transformed. Calls repositories and domain logic. |
| **Repository** | Data persistence and third-party integration. Implements storage logic (PostgreSQL, etc.). |
| **Delivery (REST)** | Handles HTTP requests and responses using Echo. Maps JSON ↔ DTO ↔ Entities. |
| **Config**   | Dependency wiring (DI), configuration loading, and server setup. |
| **DB Migrate** | Database version control using SQL migrations. |

---

### 🐳 Docker-Based Deployment

- Uses `Dockerfile` in `build` for containerizing the account service.
- Managed by `docker-compose.yaml`, which defines two main services:
  - `account-service`: The account-domain-service (REST API).
  - `postgres-db`: The PostgreSQL database.

> Database service and API container are in the same Docker network, allowing the API to resolve the DB hostname directly using `postgres-db:5432`.

---

# Development Guide

## Introduction

Welcome to the **Account Domain Service** project! This guide will walk you through the steps to set up the development environment, install dependencies, and run the project locally. It also explains how to use the provided Makefile to automate common tasks such as building, migrating databases, and generating code.

## Prerequisites

Before you can start working on this project, ensure that you have the following software installed:

- **Go (version 1.24 or above)**: Go is the primary language used for this project. You can install Go by following the [official installation guide](https://golang.org/doc/install).
- **Docker**: Docker is used to containerize the application. Install it by following the [Docker documentation](https://www.docker.com/get-started).
- **Make**: Make is used for automating tasks. It is optional but recommended for running the commands in the Makefile. Install Make via [GNU Make](https://www.gnu.org/software/make/).



## 1. Clone Project
```bash
git clone https://github.com/imansohibul/account-domain-service.git
cd account-domainn-service
```

## 2. Configure Environment
```bash
cp .env.sample .env
nano .env  # Edit with your configuration
```


## 3. Running the Service
### Option A: With Docker
```bash
cp .env.sample .env
nano .env  # Edit with your configuration
```

```bash
# Start all services (app + PostgreSQL)
docker-compose up -d --build

# View logs
docker-compose logs -f account-service
```
### Option B: Without Docker

```bash
# First start PostgreSQL container
docker-compose up -d postgres-db

# Then run the service
make run
```

## 4. Database Migrations
### Create New Migration
```bash
make create-db-migration MIGRATE_NAME=init_schema
```
This creates new files in db/migrate/

### Run Migrations
```bash
make migrate up
```

### Rollback Migrations
```bash
make migrate down N=1  # Rollback 1 step
```

## 5. Generate Mock
```bash
make generate
```
Generates mock files using mockgen


## 6. Common Commands

| Command                  | Description                              | Example Usage                     |
|--------------------------|------------------------------------------|-----------------------------------|
| `make build`             | Build the application binary             | `make build`                     |
| `make run`               | Run the service locally                  | `make run`                       |
| `make format`            | Format all Go code                       | `make format`                    |
| `make download`          | Download Go dependencies                 | `make download`                  |
| `make generate`          | Generate mock files                      | `make generate`                  |
| `make migrate up`        | Apply all pending migrations             | `make migrate up`                |
| `make migrate down`      | Rollback migrations                      | `make migrate down N=1`          |
| `make migrate status`    | Check migration status                   | `make migrate status`            |
| `make create-db-migration` | Create new migration file            | `make create-db-migration NAME=create_users` |
| `docker-compose up`      | Start all services with Docker           | `docker-compose up -d`           |
| `docker-compose logs`    | View service logs                        | `docker-compose logs -f account-service` |

**Key Flags:**
- `N=1` - Specifies number of migrations to rollback
- `NAME=migration_name` - Sets name for new migrations
- `-d` - Run Docker containers in detached mode
