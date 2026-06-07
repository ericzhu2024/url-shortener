# URL Shortener

A full-stack URL shortening service built with Go and React.

## Features

- Shorten long URLs into compact links
- Redirect to original URLs via short codes
- Click tracking and statistics
- Redis caching for fast redirects
- Rate limiting (10 requests per minute per IP)
- Duplicate URL detection

## Tech Stack

**Backend**
- Go + Gin framework
- PostgreSQL + GORM
- Redis (caching & rate limiting)

**Frontend**
- React + TypeScript
- Tailwind CSS
- Vite

**Infrastructure**
- Docker + Docker Compose
- AWS EC2

## Architecture
Browser → React Frontend (Nginx)
↓
Go Backend (Gin)
↓
PostgreSQL + Redis

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Run Locally

```bash
git clone https://github.com/ericzhu2024/url-shortener.git
cd url-shortener
docker compose up --build
```

Visit `http://localhost` for the frontend and `http://localhost:8080` for the API.

### API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/api/links` | Create a short link |
| GET | `/r/:code` | Redirect to original URL |
| GET | `/api/links/:code/stats` | Get link statistics |

## How It Works

1. User submits a long URL
2. Backend checks Redis cache for existing short code
3. If not cached, checks PostgreSQL database
4. If not found, generates a new 6-character short code
5. Stores in PostgreSQL and caches in Redis
6. Returns the short URL to the user

On redirect, Redis is checked first to minimize database load. Click counts are updated asynchronously to avoid slowing down redirects.

## Running Tests

```bash
cd backend
go test ./...
```