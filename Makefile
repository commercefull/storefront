.PHONY: start fmt lint test tidy

start:
	nodemon -e go,json,html --exec go run ./cmd/serve/main.go --signal SIGTERM

fmt:
	go fmt ./...

lint:
	go vet ./...

test:
	go test ./...

tidy:
	go mod tidy
