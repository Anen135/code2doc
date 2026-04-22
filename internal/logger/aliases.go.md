# Technical Documentation: Logger Aliases Package

## Overview

The `aliases.go` file located in the `internal/logger` package provides a set of convenient shorthand functions for logging messages at different severity levels. These functions serve as aliases to simplify the logging process across the application, offering a consistent and easy-to-use interface for developers.

---

## File Location

```
internal/logger/aliases.go
```

This file is part of the internal logging system and is not intended for external use outside the module.

---

## Package Dependencies

Before diving into individual functions, this package relies on the core logger implementation (likely defined elsewhere in the same module) that handles the actual message output, formatting, and severity filtering. These alias functions typically call into that core system.

---

## Functions

### 1. `Info(msg string)`

Logs an informational message. This level is used for general operational messages that highlight the progress of the application at a coarse-grained level.

- **Usage Example:**
  ```go
  logger.Info("Application started successfully")
  ```

- **Common Use Cases:**
  - Service initialization
  - Routine status updates
  - Non-critical events

---

### 2. `Error(msg string)`

Logs an error message. This level indicates a significant problem that has affected part of the application flow. Errors are typically unexpected and may require attention.

- **Usage Example:**
  ```go
  logger.Error("failed to connect to the database")
  ```

- **Common Use Cases:**
  - Failed operations
  - Critical system failures
  - Recoverable or non-recoverable errors during execution

---

### 3. `ErrorV(msg string, v ...interface{})`

Logs an error message with variable arguments, allowing formatted output similar to `fmt.Printf`.

- **Parameters:**
  - `msg` (string): The format string.
  - `v` (`...interface{}`): Optional arguments to format into the message.

- **Usage Example:**
  ```go
  logger.ErrorV("failed to process request for user %s (ID: %d)", "Alice", 123)
  ```

- **Common Use Cases:**
  - Detailed error reporting with context
  - Structured logging with placeholders

---

### 4. `Warning(msg string)`

Logs a warning message. Warnings indicate potential issues that are not immediately problematic but may lead to errors or unexpected behavior if left unaddressed.

- **Usage Example:**
  ```go
  logger.Warning("deprecated configuration option detected")
  ```

- **Common Use Cases:**
  - Misconfigurations
  - Deprecation notices
  - Performance or usage warnings

---

### 5. `Debug(msg string)`

Logs a debug-level message. This level is typically used for diagnostic information useful to developers during troubleshooting. Debug logs are often disabled in production environments.

- **Usage Example:**
  ```go
  logger.Debug("entering data processing pipeline with payload size 2048")
  ```

- **Common Use Cases:**
  - Detailed flow tracing
  - Variable inspection during development
  - Low-level system diagnostics

---

## Design Notes

- These alias functions are intended to provide a clean and consistent API across the application.
- The actual implementation of the logging logic (e.g., output formatting, log level filtering, destination handling) is abstracted elsewhere and is invoked by these functions.
- Variable argument support (`ErrorV`) enables structured and safe message formatting without requiring manual concatenation.

---

## Summary Table

| Function       | Level    | Description                            | Variants Supported |
|----------------|----------|----------------------------------------|---------------------|
| `Info`         | Info     | General operational messages           | No                  |
| `Error`        | Error    | Significant problems                   | No                  |
| `ErrorV`       | Error    | Formatted error messages               | Yes                 |
| `Warning`      | Warning  | Potential issues                       | No                  |
| `Debug`        | Debug    | Diagnostic information (dev use only)  | No                  |

---

## Usage Guidelines

- Use `Info` for general status updates.
- Use `Error` or `ErrorV` when failures occur.
- Use `Warning` for non-critical anomalies.
- Use `Debug` sparingly and only during development or troubleshooting.

Ensure that log levels are appropriately configured in the environment to avoid excessive output in production (e.g., disable `Debug` logs).