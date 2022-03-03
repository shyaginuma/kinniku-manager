.PHONY: build
build:
	docker-compose build

.PHONY: run-server
run-server:
	docker-compose up backend --build

.PHONY: go-test
go-test:
	docker-compose up backend-test --build
