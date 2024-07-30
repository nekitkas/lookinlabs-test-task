# Go Template Repository

Go template repository used for services written in Go

Available Make functions:

- make build - Compile the code
- make air - Run the code with Golang Hot Reloader (Air)
- make run - Run the code without Hot Reloader
- make fumpt - Run gofumpt
- make mod-vendor - Run go mod vendor
- make linter - Run golangci-lint
- make gosec - Run gosec
- make test - Run tests
- make validate - Run all linter, gosec and tests
- make migrate-up - Run database migrations
- make migrate-down - Rollback database migrations
- make migrate-create - Create a new database migration

