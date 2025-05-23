# Coptic Scribe - Dictionary Backend

This is the backend API for the **Coptic Scribe** project, a Coptic–Arabic–English dictionary that allows fast search, filtering, and contextual analysis of Coptic words. The backend is written in **Go (Gin + GORM)** and uses **PostgreSQL** as the database.

It includes a Python script to generate and populate the database from an Excel source, and supports Docker-based deployment for easy local development.

---

## ✨ Features

- ✅ RESTful API to retrieve and search Coptic dictionary entries  
- ✅ Auto-creates the database from Excel with one command  
- ✅ Fully containerized Go backend via Docker  
- ✅ CLI prompts for DB credentials if needed  
- ✅ No external database install required

---

## 🚀 Quick Start

### 1. Clone the repo

```bash
git clone https://github.com/haddad-github/coptic-scribe-dictionary-backend.git
cd coptic-scribe-dictionary-backend
```

### 2. Create `secrets.json` for your local DB config

You’ll be prompted for database name, user, and password:

```bash
make generate-secrets
```

This creates both:

- `secrets.json`
- `api/.env` (used by Docker)

---

### 3. Install Python dependencies

You need this to populate the database:

```bash
make setup
```

---

### 4. Create and seed the PostgreSQL database

Make sure your DB is running (e.g. PostgreSQL is installed or Dockerized), then run:

```bash
make create-db
```

This will:

- Create the database if it doesn't exist  
- Create the required table  
- Load entries from your Excel file  
- Add an index on `coptic_word`

---

### 5. Build and run the backend API using Docker

```bash
make run-backend
```

This will:

- Build the Go backend with Docker  
- Run it on `http://localhost:8080`

---

## 🧪 API Endpoints

| Method | Route       | Description                        |
|--------|-------------|------------------------------------|
| GET    | `/words`    | List all Coptic dictionary entries |
| GET    | `/word?id=` | Fetch a specific entry by ID       |

More endpoints coming soon.

---

## 🔄 All-in-One Setup

Want to do everything in one shot?

```bash
make run
```

This will:

- Clean any running containers  
- Prompt for credentials  
- Generate secrets  
- Setup Python deps  
- Create the database  
- Build & run the Go backend

---

## 🧼 Clean Up

Stop and remove all backend-related containers:

```bash
make clean
```

---

## 🛠️ Requirements

- [Go](https://golang.org/dl/)
- [Docker](https://www.docker.com/)
- [Python 3.9+](https://www.python.org/downloads/)
- pip + venv recommended

---

## 🤝 Contributing

Want to improve the backend or the dataset? Fork the repo, create a feature branch, and submit a PR!

---

## 📜 License

MIT License © 2025 – Rafic Haddad
