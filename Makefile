.PHONY: test tidy bench

tidy:
	@echo "Tidying up module..."
	go mod tidy

test:
	@echo "Running tests..."
	go test ./...

bench:
	@echo "Running benchmarks..."
	go test -run='^$$' -bench=. -benchmem .