test:
	go test
	go test ./cmd/pgo
	go test ./cmd/pgo-discord
install:
	go install ./cmd/pgo
	go install ./cmd/pgo-discord
build:
	go build -o=./bin/ ./cmd/pgo
	go build -o=./bin/ ./cmd/pgo-discord
format:
	gofmt -w .
	goimports -w .
	go mod tidy
