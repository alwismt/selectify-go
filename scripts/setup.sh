#!/bin/bash

# Create the main project directory
mkdir -p cmd/admin internal/admin-app/config internal/admin-app/domains internal/admin-app/interfaces/controllers/api internal/admin-app/interfaces/controllers/web internal/admin-app/interfaces/middleware internal/entities internal/infrastructure/persistence/postgres internal/infrastructure/web web/static web/templates

# Create the necessary files
touch cmd/admin/main.go internal/admin-app/config/config.go internal/admin-app/interfaces/middleware/auth_middleware.go internal/admin-app/interfaces/middleware/site_middleware.go internal/admin-app/interfaces/middleware/customer_middleware.go internal/admin-app/interfaces/middleware/product_middleware.go internal/infrastructure/persistence/postgres/pgsql.go internal/infrastructure/web/server.go .env.example .gitignore README.md .env

# Add the necessary package declarations to the files
echo "package main" > cmd/admin/main.go
echo "package config" > internal/admin-app/config/config.go
echo "package middleware" > internal/admin-app/interfaces/middleware/auth_middleware.go
echo "package middleware" > internal/admin-app/interfaces/middleware/site_middleware.go
echo "package middleware" > internal/admin-app/interfaces/middleware/customer_middleware.go
echo "package middleware" > internal/admin-app/interfaces/middleware/product_middleware.go
echo "package postgres" > internal/infrastructure/persistence/postgres/pgsql.go
echo "package web" > internal/infrastructure/web/server.go

# Output the folder structure
echo "Folder structure created"
