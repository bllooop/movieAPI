build:
		docker compose build movieapi
run:
		docker compose up movieapi
migrate:	
		migrate -path ./schema -database 'postgres://postgres:54321@localhost:5436/postgres?sslmode=disable' up
migrate-down:	
		migrate -path ./schema -database 'postgres://postgres:54321@localhost:5436/postgres?sslmode=disable' down
test:
	go test -v ./...
swag:
	swag init -g cmd/main.go