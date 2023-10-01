.PHONY: test tidy

tidy:
	@echo "Tidying up module..."
	go mod tidy

test:
	@echo "Running tests..."
	go test ./...