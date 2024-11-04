# BreizhSport Backend API

This is the backend API for BreizhSport, a sports equipment e-commerce platform.

## Table of Contents

1. [Project Structure](#project-structure)
2. [Technologies Used](#technologies-used)
3. [Prerequisites](#prerequisites)
4. [Getting Started](#getting-started)
5. [API Endpoints](#api-endpoints)
6. [Making Requests](#making-requests)
7. [Development](#development)
8. [Contributing](#contributing)
9. [License](#license)

## Project Structure

The project follows a typical Go application structure that you can find in [arch.md](./app/docs/arch.md) file.

## Technologies Used

- Go 1.23
- Gin Web Framework
- GORM (Go Object Relational Mapper)
- PostgreSQL
- Docker & Docker Compose

## Prerequisites

- Docker
- Docker Compose
- Go 1.23 (for local development)

## Getting Started

1. Clone the repository:

   ```
   git clone https://github.com/PrismeDroiteExt/bzhspback.git
   cd breizhsport-backend
   ```

2. Create a `.env` file in the root directory. You can use the `env-create.sh` script to help you set up the environment variables:

   ```
   ./env-create.sh
   ```

3. Build and start the Docker containers:
   ```
   docker-compose up --build
   ```

The API should now be running and accessible at `http://localhost:8080`.

## API Endpoints

The main API routes are defined in `app/internal/api/v1/routes.go`. Here are the current endpoints:

- Products:

  - GET `/api/v1/products/:id`: Get a product by ID
  - GET `/api/v1/products/category/:category_id`: Get products by category ID

- Brands:
  - GET `/api/v1/brands`: Get all brands
  - GET `/api/v1/brands/:id`: Get a brand by ID

## Making Requests

You can use tools like cURL, Postman, or any HTTP client to make requests to the API. Here are some examples:

1. Get a product by ID:

   ```
   curl http://localhost:8080/api/v1/products/1
   ```

2. Get products by category ID:

   ```
   curl http://localhost:8080/api/v1/products/category/1
   ```

3. Get products by category ID with filters:
   ```
   curl "http://localhost:8080/api/v1/products/category/1?brand=1,2,3&min_price=50&max_price=100&colors=red,blue&sizes=M,L"
   ```

## Development

The application uses Go Modules for dependency management. To add or update dependencies, use the `go get` command.

The `docker-compose.yml` file is set up for local development, with hot-reloading enabled. Any changes you make to the Go files will trigger a rebuild of the application.

## License

This project is licensed under the Apache License 2.0. See the [LICENSE](LICENSE) file for details.
