# go-chi-subscription-manager

## Features

- CRUD operations for subscriptions
- Calculation of total subscription cost for a period
- Swagger documentation
- Docker containerization

## Project Structure

```
.
├── app/                # Application initialization and routing
├── cmd/server/         # Entry point
├── docs/               # Swagger documentation
├── internal/           # Internal packages (common, subscription)
├── .env.example
├── .gitignore
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
```

## Getting Started

### Requirements

- [Go 1.24+](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Docker](https://www.docker.com/products/docker-desktop/)

### Configuration

Copy `.env.example` to `.env` and fill in the environment variables:

```sh
cp .env.example .env
```

### Running with Docker

Build and start the application and PostgreSQL database:

```sh
docker-compose up --build
```

The API will be available at `http://localhost:7070`.

### Running Locally

1. Start the PostgreSQL database
2. Set up the `.env` file.
3. Run the server:

```sh
go run /cmd/server/main.go
```

## API Documentation

Swagger UI is available at:  
[http://localhost:7070/swagger/index.html](http://localhost:7070/swagger/index.html)

## List of Endpoints

- `GET /api/subscriptions?user-id={uuid}` — List user subscriptions
- `POST /api/subscriptions` — Create a subscription
- `GET /api/subscriptions/{id}` — Get subscription by ID
- `PUT /api/subscriptions/{id}` — Update a subscription
- `DELETE /api/subscriptions/{id}` — Delete a subscription
- `GET /api/subscriptions/total-price?user-id={uuid}&service-name={name}&from=MM-YYYY&to=MM-YYYY` — Calculate total