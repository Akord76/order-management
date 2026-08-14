# Order Management API (Go)

A layered Go REST API (Gin + GORM + SQL Server) implementing the
`OrderManagement` schema, with JWT authentication and role-based
authorization.

## Architecture

```
Request
  │
  ▼
Router (routes/routes.go)
  │
  ▼
JWT Middleware (middleware/jwt.go)   -- validates Bearer token, sets user identity in context
  │
  ▼
Role Authorization (middleware/authorization.go) -- RequireRoles(...)
  │
  ▼
Handler (handler/*.go)   -- parses request, calls service
  │
  ▼
Service (service/*.go)   -- validation & business rules
  │
  ▼
Repository (repository/*.go) -- GORM queries / stored procedures
  │
  ▼
SQL Server (OrderManagement DB)
```

## Role hierarchy

```
                     JWT Authentication
                           │
                           ▼
                    JWT Middleware
                           │
                           ▼
                    User Identity
                           │
                           ▼
                    Role Authorization
                    /       |       \
                   /        |        \
                ADMIN    MANAGER     USER
```

| Role    | Permissions                                             |
|---------|----------------------------------------------------------|
| ADMIN   | Create User · Update User · Delete User · View Users     |
| MANAGER | Create Order · Update Order · View Order                 |
| USER    | View Order                                                |

Enforced in `routes/routes.go`:

- `/api/users/**` → `middleware.RequireRoles(middleware.RoleAdmin)`
- `POST/PUT/DELETE /api/orders/**` → `RequireRoles(RoleAdmin, RoleManager)`
- `GET /api/orders/**` → `RequireRoles(RoleAdmin, RoleManager, RoleUser)`
- Same read/write split applied to `/api/order-sales/**`

Role names are defined once as constants in `middleware/authorization.go`
(`RoleAdmin`, `RoleManager`, `RoleUser`) so routes never hardcode strings.

## Auth flow

1. `POST /api/auth/register` — self-service signup (defaults to `USER` role)
2. `POST /api/auth/login` — verifies password with bcrypt, returns a signed JWT
3. Every protected request sends `Authorization: Bearer <token>`
4. `middleware.JWTAuth` parses the token and stores `userID`, `username`,
   `role` in the Gin context
5. `middleware.RequireRoles` reads the role from context and allows/denies

An ADMIN can additionally provision accounts with any role via
`POST /api/users` (see `handler/user_handler.go` / `service/user_service.go`),
separate from the public self-registration endpoint.

## Setup

```bash
cp .env.example .env      # edit DB credentials / JWT secret
go mod tidy                # fetches gin, gorm, jwt, bcrypt, godotenv, sqlserver driver
go run main.go
```

The database, tables, and stored procedures (`usp_AppUser_CRUD`,
`usp_AppUser_Login`) are expected to already exist — run the provided
SQL script against SQL Server first.

## Endpoints (summary)

| Method | Path                                   | Roles                     |
|--------|-----------------------------------------|----------------------------|
| POST   | /api/auth/register                      | public                     |
| POST   | /api/auth/login                         | public                     |
| GET    | /api/users                               | ADMIN                      |
| POST   | /api/users                               | ADMIN                      |
| PUT    | /api/users/:id                           | ADMIN                      |
| DELETE | /api/users/:id                           | ADMIN                      |
| GET    | /api/orders                              | ADMIN, MANAGER, USER       |
| POST   | /api/orders                              | ADMIN, MANAGER             |
| PUT    | /api/orders/:orderID/:orderNo            | ADMIN, MANAGER             |
| DELETE | /api/orders/:orderID/:orderNo            | ADMIN, MANAGER             |
| GET    | /api/order-sales                         | ADMIN, MANAGER, USER       |
| POST   | /api/order-sales                         | ADMIN, MANAGER             |
| *      | /api/categories, /products, /customers,  | any authenticated user     |
|        | /employees, /suppliers                   |                            |
