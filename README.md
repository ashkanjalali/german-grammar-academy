# German Grammar Academy (A1–C2)

A full-stack starter application for learning German grammar from CEFR A1 through C2.

## Stack
- Go 1.27, Gin, GORM, PostgreSQL
- Angular 22 standalone components, TypeScript
- Docker Compose

## Features in this starter
- CEFR level browser (A1, A2, B1, B2, C1, C2)
- Grammar lessons with explanations, examples and English translations
- Practice questions with instant feedback
- Timed-style level exam endpoint and automatic multiple-choice grading
- Progress summary endpoint
- PostgreSQL persistence for lessons, questions, attempts and progress
- Seeded starter content for every CEFR level

> Content note: this is a working starter with representative lessons and question samples, not a complete accredited A1–C2 textbook or official Goethe/telc exam. Expand the curriculum and have advanced C1/C2 material reviewed by qualified German teachers before relying on it for certification preparation.

## Run with Docker
```bash
docker compose up --build
```
Then open:
- Angular: http://localhost:4200
- API health: http://localhost:8080/api/health

The Angular dev server proxies `/api` to the Go API.

## Run locally
1. Start PostgreSQL and create a database named `german_academy`.
2. Copy `backend/.env.example` to `backend/.env` and set `DATABASE_URL`.
3. Backend:
   ```bash
   cd backend
   go mod tidy
   go run ./cmd/api
   ```
4. Frontend:
   ```bash
   cd frontend
   npm install
   npm start
   ```

## Main API
- `GET /api/health`
- `GET /api/levels`
- `GET /api/lessons?level=A1`
- `GET /api/lessons/:id`
- `GET /api/exams/:level`
- `POST /api/exams/:level/submit` body: `{"answers":{"1":0,"2":2}}`
- `GET /api/progress?userId=demo`

Question answer indexes are zero-based. The demo app uses `userId=demo`; replace this with authenticated user identity before production.

## Production checklist
- Add registration/login, password hashing and short-lived access tokens + refresh-token rotation.
- Authorize every progress/attempt query by the authenticated user; do not trust a client-supplied user ID.
- Add migrations, rate limiting, structured logs, request IDs, backups and observability.
- Add richer exam sections (reading, listening, writing and speaking), rubrics and human/AI-assisted writing feedback.
- Add editorial workflow and source/reviewer metadata for curriculum content.
- Add accessibility, localization, automated tests and CI/CD.


© Copyright & Intellectual Property Notice
Let's Learn Deutsch is an original project created and owned by Ashkan Jalali.
All rights reserved. The source code, original design, branding, documentation, graphics, and other original materials are protected by applicable intellectual property laws.
No one may copy, reproduce, redistribute, modify, publish, or commercially use this project's original materials without prior written permission from the owner.
This project is proprietary. Viewing the repository does not grant permission to reuse or distribute its contents.
Please respect the original work, creativity, and development behind Let's Learn Deutsch.
© 2026 Ashkan Jalali. All Rights Reserved.
