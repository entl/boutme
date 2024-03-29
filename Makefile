-include .env
export

build_run:
	@echo "Building and running the project..."
	@go build -o bin/main cmd/boutme/main.go
	@./bin/main

migrate_up:
	@echo "Migrating up..."
	@goose --dir sql/schema $(MIGRATION_DB_USER) $(MIGRATION_DB_URL) up

migrate_down:
	@echo "Migrating down..."
	@goose --dir sql/schema $(MIGRATION_DB_USER) $(MIGRATION_DB_URL) down

