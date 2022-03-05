.PHONY: build
build:
	docker-compose build

.PHONY: run-server
run-server:
	docker-compose up --build backend

.PHONY: go-test
go-test:
	docker-compose up db -d
	docker-compose up --build --remove-orphans --exit-code-from backend_test backend_test
