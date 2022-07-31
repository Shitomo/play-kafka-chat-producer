export ENV := local
export PORT := 8080

run:
	-@rm wire_gen.go
	@go generate ./...
	@go run cmd/chat/main.go cmd/chat/http.go
lint:
	@docker run --rm -v $(pwd):/app -w /app golangci/golangci-lint:v1.47.2 golangci-lint run -v