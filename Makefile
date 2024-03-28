build_run:
	@echo "Building and running the project..."
	@go build -o bin/main cmd/boutme/main.go
	@./bin/main
