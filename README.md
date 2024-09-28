# Imagera - Image Upload and Download API

## Overview
**Imagera** is an API for uploading and downloading generative AI images. This project uses **Golang** for the backend, **Gorm** for ORM, **PostgreSQL** for the database, and **Docker** for the local environment.

## Requirements

- Docker
- Docker Compose
- Go 1.18 or higher

## Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/imagera.git
cd imagra
```

### 2. Set up environment variables

Create a .env file in the root directory with the following content:
```
# .env

DB_HOST=postgres
DB_USER=youruser
DB_PASSWORD=yourpassword
DB_NAME=yourdb
DB_PORT=5432
```

### 3. Build and run the containers

Use Docker Compose to build the containers for PostgreSQL and the Golang application.
```
docker-compose up --build
```
This command will:

- Start the PostgreSQL database.
- Start the Go API server on port 8080.

4. Apply database migrations

The migrations are automatically applied when the application starts. You can verify the database tables are created by connecting to the PostgreSQL container.


```
docker exec -it my-postgres-db psql -U youruser -d yourdb
```

### 5. Accessing the API

The API is available at http://localhost:8080.

Example endpoints:

- Health check: GET /health
- Upload image: POST /api/upload
- List images: GET /api/images
- Download image: GET /api/images/:imageId

