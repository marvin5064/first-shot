# Makefile
APPNAME=redis-streams

.PHONY: build run run-redis run-redis-cli

build:
	@go build
run: build
	@./$(APPNAME)

run-redis:
	docker run --name redis-test-with-cli -d -p 127.0.0.1:6379:6379 redis:5.0-rc
run-redis-cli:
	docker exec -it redis-test-with-cli /bin/bash