.PHONY: build
build:
	docker-compose build

.PHONY: run-server
run-server:
	docker-compose up backend

.PHONY: go-test
go-test:
	docker-compose up backend-test
