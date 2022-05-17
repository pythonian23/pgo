install: format
	@echo Installing CMD
	@go install ./cmd/pgo
test: format
	@echo Testing PGO
	@go test
	@echo Testing CMD
	@go test ./cmd/pgo
build: format
	@echo Building CMD
	@go build -o=./bin/ ./cmd/pgo
format:
	@gofmt -w .
	@goimports -w .
	@go mod tidy
