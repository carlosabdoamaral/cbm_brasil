api_token:
	go run cmd/api/main.go -token true

api:
	go run cmd/api/main.go

rabbit:
	go run cmd/rabbit/main.go

grpc:
	go run cmd/grpc/main.go

infra:
	cd deployment/ && docker-compose up -d

migrateup:
	migrate -path ./internal/database/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path ./internal/database/migrations -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

# createMigration:
# migrate create -ext sql -dir internal/database/migrations -seq init_schema