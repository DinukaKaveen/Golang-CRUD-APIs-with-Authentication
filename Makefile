include .env
export

DB_URL=mysql://$(DB_USERNAME):$(DB_PASSWORD)@tcp($(DB_HOST):$(DB_PORT))/$(DB_DATABASE)
MIGRATIONS_DIR=internal/database/migrations

sqlc-generate:
	cd internal/database && sqlc generate

migrate-create:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq create_books_table

migrate-up:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) up

migrate-down:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) down 1

migrate-down-all:
	migrate -database "$(DB_URL)" -path $(MIGRATIONS_DIR) down