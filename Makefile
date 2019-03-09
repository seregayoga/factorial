.PHONY: all
all: build

.PHONY: build
build:
	docker-compose build

.PHONY: up
up: build
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.PHONY: test
test: build
	docker run --rm factorial_factorial go test -v -cover ./...