build:
	go build -o ./cmd/short-url ./internal/main.go

execute:
	./cmd/short-url

run:
	go run ./internal/main.go