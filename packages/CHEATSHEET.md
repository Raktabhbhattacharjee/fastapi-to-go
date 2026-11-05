# Golang Packages, Modules & Visibility Cheatsheet

## 1. Project Initialization & Lifecycle (Python vs Go)

| Lifecycle Stage | Python (`uv`) | Go (`go mod`) |
| :--- | :--- | :--- |
| **1. Create directory** | `mkdir my_app && cd my_app` | `mkdir my_app && cd my_app` |
| **2. Initialize config** | `uv init` *(creates `pyproject.toml`)* | **`go mod init my_app`** *(creates `go.mod`)* |
| **3. Virtual environment**| `uv venv` *(Needed to avoid clashes)* | **NOT NEEDED AT ALL!** (Isolated by `go.mod`) |
| **4. Add dependencies** | `uv add gin` | `go get github.com/...` |
| **5. Lockfile** | `uv.lock` | **`go.sum`** *(Cryptographic SHA-256 hashes)* |
| **6. Clean dependencies**| `uv lock` | **`go mod tidy`** *(Auto-syncs imports)* |
| **7. Run app** | `uv run main.py` | `go run main.go` |

---

## 2. The Golden Rule of Visibility (Public vs Private)

In Go, there are **no** `public`, `private`, or `protected` keywords. Visibility is decided purely by the **first letter**:

```
Capital First Letter (A-Z)  ──► EXPORTED (Public: accessible across all packages)
lowercase First Letter (a-z) ──► UNEXPORTED (Private: accessible ONLY in same folder)
```

### Applied Across Go Language Elements:

```go
package auth

// 1. STRUCT & FIELDS
type User struct {
    Username string // PUBLIC: other packages can read/write user.Username
    Email    string // PUBLIC: other packages can read/write user.Email
    password string // PRIVATE: hidden! Compiler blocks outside access
}

// 2. FUNCTIONS
func RegisterUser(...) // PUBLIC: callable via auth.RegisterUser()
func hashPassword(...) // PRIVATE: only callable within auth/ folder

// 3. METHODS
func (u *User) VerifyPassword(...) // PUBLIC method
func (u *User) encryptField(...)   // PRIVATE method

// 4. CONSTANTS & VARIABLES
const MaxRetries = 3               // PUBLIC constant
const secretKey = "xyz123"         // PRIVATE constant
```

---

## 3. Same Folder vs Different Folder (The Golden Rule)

```
my_app/
├── main.go               <-- package main
└── auth/
    ├── user.go           <-- package auth
    └── hasher.go         <-- package auth
```

### Rule A: Inside the SAME Folder (`auth/`)
* **NEVER `import`!** 
* Files in the same folder that share `package auth` are already merged in the compiler's eyes.
* `user.go` can call private functions in `hasher.go` **directly with zero imports**:
  ```go
  // user.go
  hashPassword(raw) // Works directly! No import needed!
  ```

### Rule B: Across DIFFERENT Folders (`main.go` calling `auth/`)
* **MUST `import` the folder path!**
  ```go
  // main.go
  import "my_app/auth" // module_name/folder_name

  auth.RegisterUser(...) // Must prefix with package name
  ```

---

## 4. WHAT TO DO vs WHAT NOT TO DO

| Action | ❌ What NOT To Do | ✅ What To Do |
| :--- | :--- | :--- |
| **Calling code in same folder** | `import "./hasher"` or `import "my_app/auth"` | **Zero imports!** Call functions directly. |
| **Importing an external package**| `import "my_app/auth/user.go"` (file name) | `import "my_app/auth"` (folder name). |
| **Folder package declaration** | Mixing names: file 1 `package a`, file 2 `package b` | Every file in the folder **must declare the same package name**. |
| **Accessing private fields** | Trying to read `user.password` outside `auth` | Use public methods: `user.VerifyPassword()`. |

---

## 5. What Should & Should NOT Be in `.gitignore`

```
ALWAYS COMMIT (Track in Git):
├── *.go source files
├── go.mod        (Project root & dependency declarations)
└── go.sum        (Cryptographic integrity lockfile — NEVER ignore!)

NEVER COMMIT (Add to .gitignore):
├── *.exe, *.dll, *.so, binaries, bin/, dist/
├── .env, *.env.local (Secrets & API keys)
├── .vscode/, .idea/ (Editor/IDE configs)
└── *.out, coverage.txt (Test coverage files)
```
