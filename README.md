# Pulse

> A production-grade uptime monitor built from scratch in Go — my 12-week journey from zero to deployed.

---

## What is Pulse?

**Pulse** checks whether your websites are up or down, records latency, sends alerts when something breaks, and shows history through a live dashboard. It's not a tutorial toy — it's a real backend with authentication, a PostgreSQL database, a REST API, and a React frontend.

The twist: I'm building it as a structured learning plan, one sprint at a time, with the goal of becoming production-ready in Go within 12 weeks (~144 hours of study).

---

## The Plan at a Glance

| Phase | Title | Weeks | Outcome |
|-------|-------|-------|---------|
| 0 | Foundations | Week 1 | `go run . https://google.com` prints `UP 200 142ms`; single `main.go` |
| 1 | Concurrent CLI | Weeks 2–3 | All sites checked concurrently, repeating on an interval, with tests and graceful shutdown |
| 2 | HTTP API | Weeks 4–5 | Full REST API: add/remove monitors, query results, background scheduler |
| 3 | PostgreSQL | Weeks 6–7 | Data survives restarts; uptime %, p95 latency, indexed queries, Testcontainers |
| 4 | Production Features | Weeks 8–9 | Auth, Discord alerts, profiling, rate limiting, structured logging |
| 5 | Frontend | Weeks 10–11 | React dashboard with live SSE updates and a public status page |
| 6 | Ship It | Week 12 | Live on the internet, CI with GitHub Actions, Pulse monitors itself |

**32 sprints · ~144 hours · 1 deployed project**

---

## Phase Breakdown

### Phase 0 — Foundations (Week 1)
> By the end: `go run . https://google.com` prints `UP 200 142ms`, and an unreachable URL prints `DOWN` with a clear reason instead of crashing.

**Key concepts:** Go toolchain, structs & methods, error-as-values, `net/http` client, `flag` package, `defer`.

**Sprints:**
- **0.1** — Setup & first Go program (Days 1–2, ~4 hrs)
- **0.2** — HTTP requests & error handling (Days 3–4, ~4 hrs)
- **0.3** — Structs, methods & flags (Days 5–6, ~4 hrs)

**Interview questions:**
- Why does Go use error values instead of exceptions?
- What does `defer` do and when does it run?
- What is the difference between a value receiver and a pointer receiver?

---

### Phase 1 — Concurrent CLI (Weeks 2–3)
> By the end: `pulse -config sites.json -interval 30s` checks every site concurrently, repeats forever, shuts down cleanly on Ctrl+C, and has passing tests.

**Key concepts:** goroutines, `sync.WaitGroup`, `sync.Mutex`, channels, `select`, `time.Ticker`, `context.Context`, worker pools, OS signals, multi-package layout, table-driven tests, `httptest`.

**Sprints:**
- **1.1** — Slices, maps & JSON config (Days 1–2, ~4 hrs)
- **1.2** — Goroutines & WaitGroups (Days 3–4, ~4 hrs)
- **1.3** — Channels, select & tickers (Days 5–7, ~6 hrs)
- **1.4** — Context, worker pools & graceful shutdown (Days 8–10, ~6 hrs)
- **1.5** — Packages & first tests (Days 11–12, ~4 hrs)

**Interview questions:**
- What is the difference between concurrency and parallelism?
- How would you limit the number of concurrent HTTP requests?
- What happens if you send on a closed channel?
- How does `context` cancellation propagate?

---

### Phase 2 — HTTP API (Weeks 4–5)
> By the end: Pulse is a server. Add/remove monitors via REST, a background scheduler checks them, recent results are queryable. Storage lives behind an interface swappable for Postgres in Phase 3.

**Key concepts:** `http.Server`, `ServeMux`, Go 1.22 routing (`GET /monitors/{id}`), JSON encode/decode, REST design, `sync.RWMutex`, interfaces, dependency injection, middleware, `errgroup`, `httptest.NewRecorder`.

**Sprints:**
- **2.1** — Your first HTTP server (Days 1–2, ~4 hrs)
- **2.2** — JSON CRUD for monitors (Days 3–5, ~6 hrs)
- **2.3** — Interfaces & safe shared state (Days 6–7, ~4 hrs)
- **2.4** — Background scheduler & graceful shutdown (Days 8–10, ~6 hrs)
- **2.5** — Middleware, config & handler tests (Days 11–12, ~4 hrs)

**Interview questions:**
- How do you structure a Go HTTP service?
- Why depend on interfaces rather than concrete types?
- How do you shut down a server without dropping requests?
- What is middleware and how is it implemented in Go?

---

### Phase 3 — PostgreSQL (Weeks 6–7)
> By the end: Monitors and check results survive restarts. The API answers real questions: uptime %, average and p95 latency. Swapping `MemoryStore` for `PostgresStore` requires zero handler changes.

**Key concepts:** Docker Compose, schema design, SQL joins/aggregates/percentiles, migrations (`golang-migrate`), `pgx`/`pgxpool`, parameterized queries, indexes, `EXPLAIN ANALYZE`, Testcontainers.

**Sprints:**
- **3.1** — Docker, Postgres & SQL refresher (Days 1–3, ~6 hrs)
- **3.2** — Migrations (Days 4–5, ~3 hrs)
- **3.3** — Connecting Go to Postgres with pgx (Days 6–8, ~6 hrs)
- **3.4** — Queries that matter (Days 9–11, ~6 hrs)
- **3.5** — Integration tests (Days 12–13, ~4 hrs)

**Interview questions:**
- What is an index and when does it not help?
- How does connection pooling work and why does it matter?
- How do you run schema migrations safely in production?
- What is SQL injection and how do you prevent it?

---

### Phase 4 — Production Features (Weeks 8–9)
> By the end: Pulse is multi-user, sends alerts when a site goes down/recovers, handles 1,000+ monitors with stable resource use, and logs like a real service.

**Key concepts:** bcrypt, JWTs vs sessions, auth middleware, state machines (UP/DOWN incidents with thresholds), webhooks + exponential backoff, `pprof`, `x/time/rate`, `slog`, server timeouts, `golangci-lint`.

**Sprints:**
- **4.1** — Users & authentication (Days 1–3, ~6 hrs)
- **4.2** — Incidents & alerts (Days 4–6, ~6 hrs)
- **4.3** — Scaling the scheduler (Days 7–8, ~4 hrs)
- **4.4** — Logging & hardening (Days 9–11, ~6 hrs)
- **4.5** — Quality pass (Days 12–13, ~4 hrs)

**Interview questions:**
- How do you store passwords securely?
- JWT or sessions: what are the trade-offs?
- How would you design alerting that avoids false alarms?
- How do you find a memory leak in a Go service?

---

### Phase 5 — Frontend (Weeks 10–11)
> By the end: A dashboard where users sign in, manage monitors, see uptime charts with live updates, and share a public status page.

**Key concepts:** CORS, pagination, React (or HTMX), authenticated API calls from the browser, time-bucketed queries with `date_trunc`, Chart.js, Server-Sent Events, `http.Flusher`, fan-out pub/sub with channels, in-memory caching.

**Sprints:**
- **5.1** — Getting the API ready for a browser (Days 1–2, ~4 hrs)
- **5.2** — Auth pages & monitor management (Days 3–5, ~6 hrs)
- **5.3** — Monitor detail page with charts (Days 6–8, ~6 hrs)
- **5.4** — Live updates with Server-Sent Events (Days 9–10, ~4 hrs)
- **5.5** — Public status page (Days 11–12, ~4 hrs)

**Interview questions:**
- What is CORS and why does the browser enforce it?
- Polling, SSE or WebSockets: when would you pick each?
- How would you cache a hot public endpoint?

---

### Phase 6 — Ship It (Week 12)
> By the end: Pulse runs live on the internet over HTTPS, every push is tested automatically, and the README explains the project in 60 seconds. Pulse monitors itself.

**Key concepts:** multi-stage Docker builds, Docker Compose for the full stack, CI with GitHub Actions, secrets in environment variables, managed Postgres.

**Sprints:**
- **6.1** — Containerize everything (Days 1–2, ~4 hrs)
- **6.2** — Continuous integration (Days 3–4, ~3 hrs)
- **6.3** — Deploy (Days 5–6, ~4 hrs)
- **6.4** — README & demo (Day 7, ~2 hrs)

**Interview questions:**
- Walk me through how Pulse works end to end.
- What would you change to support 1 million monitors?
- How does your CI pipeline protect the main branch?

---

## Target Project Structure

```
pulse/
├── cmd/
│   ├── pulse/          # CLI entrypoint
│   └── pulse-server/   # HTTP server entrypoint
├── internal/
│   ├── checker/        # HTTP check logic, worker pool
│   ├── config/         # JSON config loading
│   ├── scheduler/      # Per-monitor interval scheduler
│   └── store/          # Store interface + Memory/Postgres implementations
├── migrations/         # SQL migration files (golang-migrate)
├── web/                # React frontend
├── docker-compose.yml
├── Dockerfile
├── Makefile
└── sites.json          # Sample config for the CLI
```

---

## How I Am Working Through This

Each sprint follows the same loop:

1. **Learn** — read the linked docs and watch the recommended talks
2. **Plan** — write down the types and functions I will build, and ask Claude to review it
3. **Build** — type every line of code myself (no copy-paste)
4. **Done** — tick every "Done when" criterion, commit, and push

Weekly reviews happen on day 7 of each week, sent to Claude with a link to the latest commit.

---

## Why Go? Why an Uptime Monitor?

Go is a strong choice for backend/infrastructure work because of its concurrency primitives, fast compile times, single binary output, and wide adoption at companies like Google, Cloudflare, Uber, and Datadog.

An uptime monitor is an ideal learning project because:
- **Phase 0** already produces something useful (a real HTTP check)
- **Concurrency is natural** — checking 100 sites one at a time is visibly slow
- **Every phase adds something real** — not just refactoring the same thing
- **By the end it is deployable**, not just a tutorial toy

---

## Resources

| Topic | Resource |
|-------|----------|
| Go basics | [A Tour of Go](https://go.dev/tour/) |
| Patterns by example | [Go by Example](https://gobyexample.com/) |
| TDD in Go | [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests) |
| Concurrency talk | [Rob Pike — Concurrency is not Parallelism](https://www.youtube.com/watch?v=oV9rvDllKEg) |
| HTTP services | [Mat Ryer — How I write HTTP services after 13 years](https://grafana.com/blog/2024/02/09/how-i-write-http-services-in-go-after-13-years/) |
| Database indexing | [Use The Index, Luke](https://use-the-index-luke.com/) |
| Security | [OWASP Cheat Sheets](https://cheatsheetseries.owasp.org/) |
| Go mistakes | [100 Go Mistakes](https://100go.co/) |

---

## Progress Tracker

The interactive progress tracker lives in [`docs/Pulse · Pranav's Go plan.html`](docs/Pulse%20%C2%B7%20Pranav%27s%20Go%20plan.html) — open it in a browser to track sprints, log study sessions, and generate weekly reviews.

---

## License

MIT
