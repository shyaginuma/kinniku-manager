.PHONY: build
build:
	docker-compose build

.PHONY: run-server
run-server:
	docker-compose up --build backend

.PHONY: go-test
go-test:
	docker-compose up --build --exit-code-from backend_test backend_test
