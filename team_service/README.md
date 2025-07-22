# Teams API

A simple RESTful API to manage teams, members, and managers using Go (Gin + GORM) and PostgreSQL.

---

## 📦 Features

* Create a team
* Add/remove members
* Add/remove managers
* Auto database migration with GORM
* Health check endpoint

---

## 📚 API Specification

### `GET /health`

Check if API is running.

**Response:**

```json
{
  "status": "ok",
  "message": "Teams API is running"
}
```

---

### `POST /teams`

Create a new team.

**Request Body:**

```json
{
  "teamName": "Engineering"
}
```

**Response:**

```json
{
  "teamId": "...",
  "teamName": "Engineering",
  "ownerId": "...",
  "createdAt": "...",
  "updatedAt": "..."
}
```

---

### `POST /teams/:teamId/members`

Add a member to a team.

**Request Body:**

```json
{
  "memberId": "uuid-of-member"
}
```

---

### `DELETE /teams/:teamId/members/:memberId`

Remove a member from a team.

---

### `POST /teams/:teamId/managers`

Add a manager to a team.

**Request Body:**

```json
{
  "managerId": "uuid-of-manager"
}
```

---

## 🚀 Getting Started

### ✅ Prerequisites

* Docker
* Docker Compose

---

### 🔧 Run the App with Docker

1. **Build and start the services:**

   ```bash
   docker-compose up --build
   ```

2. **API will be available at:**
   `http://localhost:8080`

3. **PostgreSQL will be available at:**
   `localhost:5432`
   (user: `postgres`, password: `password`, db: `teams_db`)

---

## 🛠 Project Structure

```
.
├── cmd/api/
│   ├── main.go         # Entry point
│   ├── handlers.go     # API handlers
│   ├── models.go       # GORM models
│   └── database.go     # DB connection + migration
├── Dockerfile
├── docker-compose.yml
└── README.md
```
