run:
	go run cmd/main.go

run-db:
	docker run --name postgres -e POSTGRES_PASSWORD=123321 -p 5432:5432 -d --rm postgres
migrate-up:
	migrate -path migrates -database "postgres://postgres:123321@localhost:5432/postgres?sslmode=disable" up
migrate-down:
	migrate -path ./migrates -database "postgres://postgres:123321@localhost:5432/postgres?sslmode=disable" down