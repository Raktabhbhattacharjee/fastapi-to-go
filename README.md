# 🚀 FastAPI to Go — Master Road

A hands-on, structured workspace designed to transition from Python/FastAPI and C/C++ to idiomatic Go.

---

## 📂 Project Structure

```text
fastapi-to-go/
├── 01_basics/            # Values, Variables, Zero Values, Scopes, Shadowing, iota
│   └── main.go
├── 02_control_flow/      # If-Init statement, Tagless Switch, While-style & Range loops
│   └── main.go
├── 03_data_structures/   # Fixed Arrays vs Slices (make, cap), Maps, Comma-ok idiom
│   └── main.go
├── CHEATSHEET_BASICS.md  # Comprehensive quick-reference cheatsheet
├── go.mod                # Go module definition
├── main.go               # Workspace directory guide
└── README.md             # This guide
```

---

## 🏃 Running the Lessons

Run any lesson directly from the root folder:

```bash
# Lesson 1: Variables, Types, Scopes
go run ./01_basics

# Lesson 2: If/Else, Switch, Loops
go run ./02_control_flow

# Lesson 3: Slices, Maps, Data Structures
go run ./03_data_structures
```

---

## 💡 How Go Project Structure Compares to FastAPI

### In FastAPI:
```text
my_api/
  main.py          <- imports routers
  routers/         <- users.py, auth.py
  schemas/         <- Pydantic models
  services/        <- business logic
```

### In Production Go (Standard Layout):
```text
cmd/
  api/
    main.go        <- Entrypoint: reads config, wires dependencies, starts HTTP server
internal/          <- Private code (compiler prevents external packages from importing)
  handler/         <- HTTP handlers (equivalent to FastAPI routers)
  service/         <- Business logic
  model/           <- Struct definitions (equivalent to Pydantic models)
  repository/      <- Database queries (SQL / GORM / pgx)
pkg/               <- Public helper packages (optional)
```
