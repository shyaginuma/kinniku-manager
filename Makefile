.PHONY: build
build:
	docker-compose build

.PHONY: run-server
run-server:
	docker-compose down
	docker-compose up -d db
	docker-compose up --build --remove-orphans backend

.PHONY: go-test
go-test:
	docker-compose down
	docker-compose up -d db
	docker-compose up --build --remove-orphans --exit-code-from backend_test backend_test
