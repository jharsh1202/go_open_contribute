# Makefile

.PHONY: build run docker-build docker-run

build:
	go build -o main cmd/server/main.go

run:
	go run cmd/server/main.go

docker-build:
	docker build -t open-contribute .

docker-run:
	docker-compose up --build

test:
	go test ./...

clean:
	go clean
	rm -f main
	docker-compose down -v
