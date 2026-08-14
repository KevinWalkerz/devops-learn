# devops-lab

Small Go HTTP API that stores notes in Postgres.

## Run locally

1. Create a Postgres database (any way you like).
2. Copy `.env.example` to `.env` and set `DATABASE_URL`.
3. Start the API:

```bash
go run ./cmd/api
```

4. Check it:

```bash
curl http://localhost:8080/health
curl http://localhost:8080/api/notes
curl -X POST http://localhost:8080/api/notes \
  -H 'Content-Type: application/json' \
  -d '{"text":"hello"}'
```

`GET /health` pings the database. The `notes` table is created on startup if it does not exist.
