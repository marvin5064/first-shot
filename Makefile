# Makefile

.PHONY: run-redis
run-redis:
	docker run --name redis-test-with-cli -d redis:5.0-rc
run-redis-cli:
	docker exec -it redis-test-with-cli /bin/bash