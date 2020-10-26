.PHONY: build
build:
	docker-compose build

.PHONY: run
run:
	docker-compose up -d
	./scripts/wait.sh

.PHONY: clean
clean:
	docker-compose down
