init:
	make setup
	make di
	make migrate

setup:
	./setup.sh

di:
	wire ./api

migrate:
	go run cmd/migrate/main.go

main:
	go run main.go

cov:
	go test --coverprofile=cover/c.out ./...
	go tool cover --html=cover/c.out -o cover/coverage.html
	open cover/coverage.html

cov-deploy:
	go test --coverprofile=cover/c.out ./...
	go tool cover --html=cover/c.out -o cover/coverage.html

test:
	go test ./...