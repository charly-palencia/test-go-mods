### How to run 
docker run --rm -v $(pwd):/app -w /app golang:1.22-alpine sh -c "go mod init example.com/monorepo && go mod tidy"

or

make init (not needed)

### How to build

make build-users
make build-files

### How to run

make run-users
make run-files


## Validations

make validate-size-users # should be around 4.3MB
make validate-size-files # should be around 2.4MB


## Enter

make enter-users
make enter-files
# test-go-mods
