# go-api-template

A minimal Go API template using [Chi](https://github.com/go-chi/chi).

## Features

- Chi router with structured logging and panic recovery
- Health check at `/health`
- Graceful shutdown
- Multi-stage Docker build (distroless)
- Woodpecker CI: formatting, race tests, vet, build, lint (golangci-lint), security (govulncheck)
- Docker image published to GHCR on mirrored release tags by GitHub Actions

## Repository and releases

[Forgejo](https://forgejo.tail848095.ts.net/matt-riley/go-api-template) is the
source of truth for code, issues, and pull requests (Tailscale required).
[GitHub](https://github.com/matt-riley/go-api-template) is a public Git mirror
and runs only the tag-triggered GHCR publishing workflow.

Clone from Forgejo:

```sh
git clone ssh://git@forgejo.tail848095.ts.net:2222/matt-riley/go-api-template.git
```

[Woodpecker](https://woodpecker.tail848095.ts.net/) runs `.woodpecker.yaml`
for pull requests and pushes to `main`. Merge only after its checks pass.
Keep the Go module path unchanged: existing Go imports still use the GitHub
mirror, which remains available.

Create release tags only on a successful, merged `main` commit, and push them
to Forgejo. Its asynchronous push mirror forwards tags to GitHub, where
`.github/workflows/docker.yml` builds and publishes
`ghcr.io/matt-riley/go-api-template` using GitHub's own `GITHUB_TOKEN`.
It does not rerun tests or query Woodpecker; mirroring is not a CI gate.
Never tag an untested branch. Publishing still depends on GitHub availability;
ordinary reviews and CI do not.

GitHub Renovate is paused to prevent it writing to the mirror. Existing
dependency PRs were imported; a Forgejo-aware dependency updater is not yet
configured.

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
