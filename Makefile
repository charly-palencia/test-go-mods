# Nombre de la imagen base
GO_IMAGE=golang:1.22-alpine
MODULE=example.com/monorepo

# Carpetas
USERS_CMD=./cmd/users
FILES_CMD=./cmd/files

BIN_DIR=bin

# Binarios
USERS_BIN=$(BIN_DIR)/users
FILES_BIN=$(BIN_DIR)/files

# Inicializa el módulo Go si no existe
init:
	docker run --rm -v $(PWD):/app -w /app $(GO_IMAGE) sh -c "go mod init $(MODULE) || true && go mod tidy"

build-users:
	docker build --build-arg SERVICE=users -t monorepo-users .

build-files:
	docker build --build-arg SERVICE=files -t monorepo-files .

run-users: build-users
	docker run --rm monorepo-users

run-files: build-files
	docker run --rm monorepo-files

validate-size-users:
	docker run --rm monorepo-users sh -c "ls -lh /app/service" 

validate-size-files:
	docker run --rm monorepo-files sh -c "ls -lh /app/service"

enter-users:
	docker run --rm -it monorepo-users /bin/sh

enter-files:
	docker run --rm -it monorepo-files /bin/sh