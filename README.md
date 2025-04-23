## 🏗️ Project Structure & Software Architecture

This project follows a **modular clean architecture** pattern. It ensures high maintainability, testability, and clear separation of concerns.

# Account Domain Service

![Go Version](https://img.shields.io/badge/go-1.24%2B-blue)
![PostgreSQL](https://img.shields.io/badge/PostgreSQL-17.4-blue)
![Docker](https://img.shields.io/badge/docker-compose-3.9-blue)

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
├── docker-compose.yml       # Defines services (API, DB) for deployment
├── .env.sample              # Sample environment configuration
└── Makefile                 # Common scripts for building, testing, running
---
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
- Managed by `docker-compose.yml`, which defines two main services:
  - `account-service`: The account-domain-service (REST API).
  - `postgres-db`: The PostgreSQL database.

> Database service and API container are in the same Docker network, allowing the API to resolve the DB hostname directly using `db:5432`.

---

# Development Guide

## Introduction

Welcome to the **Account Domain Service** project! This guide will walk you through the steps to set up the development environment, install dependencies, and run the project locally. It also explains how to use the provided Makefile to automate common tasks such as building, migrating databases, and generating code.

## Prerequisites

Before you can start working on this project, ensure that you have the following software installed:

- **Go (version 1.24 or above)**: Go is the primary language used for this project. You can install Go by following the [official installation guide](https://golang.org/doc/install).
- **Docker**: Docker is used to containerize the application. Install it by following the [Docker documentation](https://www.docker.com/get-started).
- **Make**: Make is used for automating tasks. It is optional but recommended for running the commands in the Makefile. Install Make via [GNU Make](https://www.gnu.org/software/make/).


