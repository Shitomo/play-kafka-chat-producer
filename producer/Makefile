export ENV := local
export PORT := 8080

run:
	@rm wire_gen.go
	@go generate ./...
	@go run main.go wire_gen.go
lint:
	@docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.47.2 golangci-lint run -v