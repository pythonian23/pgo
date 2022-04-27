test:
	@echo Testing PGO
	@go test
	@echo Testing CMD
	@go test ./cmd/pgo
	@echo Testing DISCORD
	@go test ./cmd/pgo-discord
install-cmd:
	@echo Installing CMD
	@go install ./cmd/pgo
install-discord:
	@echo Installing DISCORD
	@go install ./cmd/pgo-discord
install: install-cmd install-discord
build-cmd:
	@echo Building CMD
	@go build -o=./bin/ ./cmd/pgo
build-discord:
	@echo Building DISCORD
	@go build -o=./bin/ ./cmd/pgo-discord
build: build-cmd build-discord
format:
	@gofmt -w .
	@goimports -w .
	@go mod tidy
