<h1 align="center">
  <img src="https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go&logoColor=white" alt="Go 1.22+" />
  <br />
  URL Shortener
  <br />
  <img src="https://img.shields.io/badge/Echo-v4%2Fv5-7A8CFF" alt="Echo" />
  <img src="https://img.shields.io/badge/PostgreSQL-15-336791?logo=postgresql&logoColor=white" alt="PostgreSQL" />
  <img src="https://img.shields.io/badge/Redis-7-FF4438?logo=redis&logoColor=white" alt="Redis" />
  <img src="https://img.shields.io/badge/Docker-Compose-2496ED?logo=docker&logoColor=white" alt="Docker Compose" />
  <img src="https://img.shields.io/badge/Prometheus-Metrics-E6522C?logo=prometheus&logoColor=white" alt="Prometheus" />
  <img src="https://img.shields.io/badge/Status-Portfolio%20Project-2EA043" alt="Portfolio project" />
</h1>

## The Idea

URL Shortener is a Go-based backend service for shortening long URLs with fast redirects, asynchronous click analytics, Redis-backed caching, and optional per-link monetization via interstitial advertising. It addresses the need for compact, trackable links while keeping redirect latency low and demonstrating production-oriented backend patterns in a clean, portfolio-friendly architecture.

## Features

| Feature | Status |
| --- | --- |
| Short link creation with custom alias and TTL | ✅ |
| Redirect flow with optional interstitial ad page | ✅ |
| Redis cache-aside for low-latency redirects | ✅ |
| Asynchronous click processing through Redis Streams | ✅ |
| Aggregated analytics by country, referrer, device, and day | ✅ |
| Per-link monetization toggle with ad impression tracking | ✅ |
| Rate limiting by IP and API key | ✅ |
| Graceful shutdown for API and worker services | ✅ |
| Prometheus metrics for throughput, latency, and cache hit rate | ✅ |
| Integration testing with testcontainers-go | ✅ |
| Docker Compose and CI/CD-ready deployment setup | ✅ |

## Stack

- Language: Go 1.22+
- Web framework: Echo
- Database: PostgreSQL 15
- Cache / queue: Redis 7
- Metrics: Prometheus
- Infrastructure: Docker Compose
- Testing: Go test, testcontainers-go
- CI/CD: GitHub Actions
- Linting: golangci-lint
- Local tooling: Makefile

## Structure

```text
url-shortener/
+-- cmd/                                                  # Application entrypoints
|   +-- api/                                               # API server code
|   |   L-- main.go                                        # Starts the Echo HTTP API
|   L-- worker/                                            # Background worker code
|       L-- main.go                                        # Starts click consumer and async processing
|
+-- internal/                                              # Core business logic
|   +-- link/                                              # Link management slice
|   |   +-- handler.go                                     # HTTP handlers for link CRUD and redirect
|   |   +-- service.go                                     # Link creation, validation, and lookup logic
|   |   +-- repository.go                                  # Access layer for link persistence
|   |   +-- model.go                                       # Link model and request/response DTOs
|   |   L-- cache.go                                       # Redis cache-aside implementation
|   |
|   +-- analytics/                                         # Click analytics slice
|   |   +-- consumer.go                                    # Redis stream consumer to process click events
|   |   +-- producer.go                                    # Producer for click events
|   |   +-- service.go                                     # Aggregation and enrichment logic
|   |   +-- repository.go                                  # Analytics persistence layer
|   |   L-- model.go                                       # Analytics models and aggregated stats
|   |
|   +-- monetization/                                      # Ad layer for per-link monetization
|   |   +-- provider.go                                    # AdProvider abstraction interface
|   |   +-- adsense.go                                     # Concrete ad provider or stub implementation
|   |   +-- service.go                                     # Logic for showing and tracking interstitial ads
|   |   L-- model.go                                       # Monetization data model
|   |
|   +-- ratelimit/                                         # Abuse protection layer
|   |   L-- limiter.go                                     # Token bucket and Redis-backed rate limiting
|   |
|   L-- platform/                                          # Shared infrastructure components
|       +-- postgres/                                      # PostgreSQL setup and migrations
|       |   +-- db.go                                      # Database connection and wiring
|       |   L-- migrations/                                # Schema migration files
|       |       +-- 0001_init.up.sql                       # Initial DB schema
|       |       L-- 0001_init.down.sql                     # Rollback schema
|       +-- redis/                                        # Redis integration
|       |   L-- client.go                                 # Redis client configuration
|       +-- config/                                       # Application configuration
|       |   L-- config.go                                  # Environment-based config loader
|       +-- logger/                                       # Structured logging
|       |   L-- logger.go                                  # Logger initialization
|       L-- metrics/                                       # Prometheus instrumentation
|           L-- prometheus.go                              # Metrics registration and setup
|
+-- pkg/                                                  # Reusable packages
|   +-- shortcode/                                         # Short code generation
|   |   L-- generator.go                                   # Encodes and validates short aliases/codes
|   +-- httpresponse/                                      # Shared API helpers
|       L-- response.go                                    # Common response formatting utilities
|
+-- web/                                                  # Web assets
|   L-- templates/                                         # Template directory
|       L-- interstitial.html                              # HTML ad page with countdown and redirect
|
+-- test/                                                 # Test suite
|   +-- integration/                                       # Integration tests
|   |   +-- link_test.go                                   # Link lifecycle and redirect tests
|   |   +-- analytics_test.go                              # Analytics pipeline tests
|   |   L-- monetization_test.go                           # Monetization and interstitial tests
|   L-- mocks/                                             # Mock implementations used in tests
|
+-- deploy/                                               # Deployment and runtime config
|   +-- docker-compose.yml                                 # Local infrastructure orchestration
|   +-- Dockerfile.api                                     # API container image
|   +-- Dockerfile.worker                                  # Worker container image
|   L-- prometheus.yml                                     # Prometheus scrape configuration
|
+-- docs/                                                 # Documentation
|   L-- openapi.yaml                                       # API contract definition
|
+-- .github/workflows/ci.yml                              # CI pipeline
+-- go.mod                                               # Go module definition
+-- go.sum                                               # Dependency checksums
+-- Makefile                                             # Development commands
+-- README.md                                            # Project documentation
+-- .env.example                                         # Example environment variables
+-- .gitignore                                           # Git ignore rules
L-- LICENSE                                              # License metadata, if enabled
```

## Get Started

1. Clone the repository.

```bash
git clone https://github.com/rukiamuq-hard/URL-Shortener.git
cd URL-Shortener
```

2. Configure environment variables.

Create a `.env` file or copy the example configuration:

```bash
cp .env.example .env
```

Example:

```env
APP_PORT=8080
APP_ENV=development
POSTGRES_DSN=postgres://postgres:postgres@localhost:5432/urlshortener?sslmode=disable
REDIS_ADDR=localhost:6379
RATE_LIMITER_ENABLED=true
```

3. Start required infrastructure.

```bash
docker compose -f deploy/docker-compose.yml up -d postgres redis prometheus
```

4. Run the API service.

```bash
go run ./cmd/api
```

5. Run the background worker.

```bash
go run ./cmd/worker
```

Optional Makefile shortcuts:

```bash
make up
make test
```

## Usage

Create a short link:

```bash
curl -X POST http://localhost:8080/links \
  -H "Content-Type: application/json" \
  -d '{
    "original_url": "https://example.com/very/long/path",
    "alias": "demo",
    "ttl": 86400,
    "monetization_enabled": false
  }'
```

Redirect to the target URL:

```bash
curl -I http://localhost:8080/demo
```

Get link details:

```bash
curl http://localhost:8080/links/demo
```

Update the monetization flag:

```bash
curl -X PATCH http://localhost:8080/links/demo \
  -H "Content-Type: application/json" \
  -d '{"monetization_enabled": true}'
```

View aggregated stats:

```bash
curl http://localhost:8080/links/demo/stats
```

Retrieve Prometheus metrics:

```bash
curl http://localhost:8080/metrics
```

If a link has `monetization_enabled=true`, the service serves an interstitial HTML page with a countdown and redirect flow before sending the user to the original URL. Click events are produced asynchronously so the redirect response remains fast and non-blocking.