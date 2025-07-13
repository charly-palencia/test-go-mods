# Go Monorepo Example

## Introduction

This repository demonstrates a simple Go monorepo using Go modules. It contains
two services: `users` and `files`. Each service has its own `main.go` entry
point, business logic, and embedded data.

## Folder Structure

The folder structure is as follows:

* `cmd/`: Entry points for each service
	+ `users/`: `main.go` for the Users service
	+ `files/`: `main.go` for the Files service
* `internal/`: Business logic and utilities
	+ `users/`: Logic and embedded data for Users
		- `service.go`
		- `users_mock.go`
		- `mocks/`
			- `users.json`
	+ `files/`: Logic and embedded data for Files
		- `service.go`
		- `files_mock.go`
		- `mocks/`
			- `files.json`
	+ `shares/`: Shared code between services
		- `logger.go`: Common logging utilities
		- `reader.go`: Common JSON reading helpers
* `bin/`: (Generated) Compiled binaries
	+ `users`
	+ `files`
* `Dockerfile`: Multi-stage build for Docker
* `Makefile`: Shortcuts for building, running, validating
* `README.md`: This file

## How to Run

### Initialize

To initialize the Go module, run:

```bash
make init
```

### Build

To build the services, run:

```bash
make build-users
make build-files
```

### Run

To run the services, run:

```bash
make run-users
make run-files
``` 

### Validate

To validate the services, run:

```bash
make validate-size-users # should be around 4.3MB
make validate-size-files # should be around 2.4MB
```

### Enter

To enter the services, run:

```bash
make enter-users
make enter-files
```