# Golang Struktur

use layered architecture & dependency injection pattern

## Module

module | description
--- | ---
[godotenv](https://pkg.go.dev/github.com/joho/godotenv) | environment file .env
[echo](https://pkg.go.dev/github.com/labstack/echo/v4) | router framework
[gorm](https://pkg.go.dev/gorm.io/gorm) | Object Relational Mapper
PostgreSQL | Database

## How to run

1. Database Server PostgreSQL was Running on local/server
2. update `.env` file, and fill with correct credentials

```bash
go run main.go
```