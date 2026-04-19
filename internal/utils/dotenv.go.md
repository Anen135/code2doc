# Technical Documentation: `.env` File Loader

This document describes the internal utility package responsible for loading environment variables from a `.env` configuration file. The implementation is written in Go and provides a simple interface for managing application configuration outside of the source code.

## Package Location
- **File Path:** `internal/utils/dotenv.go`
- **Visibility:** Internal (not exposed to external packages)
- **Purpose:** Provide a centralized method to load environment variables from a `.env` file during application startup.

---

## Function: `LoadEnv`

### Signature
```go
func LoadEnv() error
```

### Description
The `LoadEnv` function reads environment variables from a `.env` file located in the application's root directory and injects them into the process environment. This allows developers to manage configuration such as API keys, database connection strings, and other runtime settings without hardcoding them into the source code.

### Behavior
- Locates the `.env` file in the current working directory.
- Parses each line of the file, ignoring comments (lines starting with `#`) and empty lines.
- Supports basic key-value syntax: `KEY=VALUE`
- Exposes the parsed variables into the environment using `os.Setenv`.
- Returns an error if the file cannot be read or if an I/O error occurs during processing.
- Does **not** override existing environment variables unless explicitly configured to do so (current behavior: skips if already set).

### Error Handling
- Returns `nil` if the `.env` file does not exist (non-fatal).
- Returns an error if the file exists but cannot be read due to permission issues or malformed content.

### Example `.env` File
```env
# Application Configuration
APP_ENV=development
APP_PORT=8080

# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=admin
DB_PASSWORD=secret123
DB_NAME=myapp_db
```

After calling `LoadEnv()`, these variables become accessible via `os.Getenv("DB_HOST")`, etc.

### Usage Example
```go
package main

import (
    "internal/utils"
    "log"
)

func main() {
    if err := utils.LoadEnv(); err != nil {
        log.Fatalf("Failed to load environment variables: %v", err)
    }

    port := utils.GetEnv("APP_PORT", "3000") // Assume GetEnv is a helper
    log.Printf("Server starting on port %s", port)
}
```

---

## Design Notes

- **Security Consideration:** Sensitive values (e.g., passwords) are loaded into memory. Ensure proper file permissions on `.env` files in production environments.
- **Scalability:** The function is designed for simplicity and is not intended to handle large-scale configuration management. For advanced needs, consider integrating with a dedicated configuration library.
- **Integration:** This function is typically called early in the `main()` function to ensure environment variables are available throughout the application lifecycle.

---

## Future Enhancements (Optional)
- Support for environment-specific files (e.g., `.env.production`)
- Variable interpolation (e.g., `DB_URL=${DB_HOST}:${DB_PORT}`)
- Override protection flag (prevent system env vars from being overwritten)
- Validation of required variables on load