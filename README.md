# go-api-template

A minimal Go API template using [Chi](https://github.com/go-chi/chi).

## Features

- Chi router with structured logging and panic recovery
- Health check at `/health`
- Graceful shutdown
- Multi-stage Docker build (distroless)
- CI via shared workflows: test, lint (golangci-lint), security (govulncheck)
- Docker image published to GHCR on tags

## Getting started

```bash
# Run locally
make dev

# Run tests
make test

# Build
make build

# Build Docker image
docker build -t go-api-template .
```

## Endpoints

| Method | Path       | Description          |
| ------ | ---------- | -------------------- |
| GET    | `/health`  | Health check         |
| GET    | `/api/v1/` | Hello World endpoint |

## Configuration

| Variable | Default | Description |
| -------- | ------- | ----------- |
| `PORT`   | `8080`  | Server port |
