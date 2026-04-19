# Technical Documentation: `internal/scanner.go`

## File Overview

**File Path:** `internal/scanner.go`

This module provides functionality for scanning project files within a directory structure. It includes mechanisms for detecting and processing files while respecting visibility rules (e.g., hidden files).

---

## Data Structures

### `ProjectFile`

**Type:** `struct`

Represents metadata and content information of a file encountered during the scan process.

| Field | Type | Description |
|-------|------|-------------|
| `Path` | `string` | Absolute or relative path to the file. |
| `Name` | `string` | Base filename (without directory path). |
| `Size` | `int64` | File size in bytes. |
| `IsHidden` | `bool` | Indicates whether the file is hidden (e.g., starts with `.` on Unix systems). |
| `ModTime` | `time.Time` | Last modification timestamp. |

---

## Functions

### `Scan(rootDir string) ([]ProjectFile, error)`

Scans the directory tree starting from `rootDir` and returns a list of `ProjectFile` entries for all encountered files.

#### Parameters:
- `rootDir` (`string`): The root directory path to begin scanning from.

#### Returns:
- `[]ProjectFile`: A slice of `ProjectFile` structs representing each file found.
- `error`: An error if the scan could not be completed (e.g., permission issues, invalid path).

#### Behavior:
- Recursively traverses directories.
- Includes all non-hidden files by default (see `isHidden`).
- Captures file metadata including size and modification time.
- Skips directories and files that are inaccessible due to permission restrictions, logging warnings as needed.

#### Example Usage:
```go
files, err := Scan("/path/to/project")
if err != nil {
    log.Fatalf("Scan failed: %v", err)
}
for _, f := range files {
    fmt.Printf("Found: %s (%d bytes)\n", f.Path, f.Size)
}
```

---

### `isHidden(name string) bool`

Determines whether a file or directory is hidden based on its name.

#### Parameters:
- `name` (`string`): The filename or directory name to evaluate.

#### Returns:
- `bool`: `true` if the name starts with a dot (`.`), indicating a hidden file/directory; otherwise `false`.

#### Notes:
- This function is used internally by `Scan` to filter out hidden files.
- Platform-specific behavior is minimized; on Unix-like systems, dotfiles are hidden; on Windows, the system hidden attribute may be considered in future extensions.

#### Example:
```go
isHidden(".git")     // returns true
isHidden("main.go")  // returns false
```

---

## Design Notes

- The scanner is designed for internal use within the project tooling and assumes trusted execution context.
- Hidden files are excluded by default to reduce noise in project analysis.
- Error handling focuses on robustness: partial results may be returned if some directories cannot be read.

---

## Future Enhancements

- Add option to include hidden files via a boolean flag in `Scan`.
- Support filtering by file extensions.
- Implement concurrency (goroutines) for improved performance on large directories.