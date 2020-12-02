.PHONY: build
build:
	docker-compose build

.PHONY: download-deps
download-deps:
	go mod vendor
	go mod tidy

.PHONY: run
run:
	docker-compose up -d
	./scripts/wait.sh
	make seed

.PHONY: lint
lint:
	golangci-lint run --fix --deadline=5m ./...

.PHONY: clean
clean:
	docker-compose down

.PHONY: seed
seed:
	docker exec acp ./authorization import \
		--sql.url 'postgres://root@crdb:26257/defaultdb?sslmode=disable' \
		--secret.key KNEcLGdDqpwrXDubqPgDSUkMMsLPXaHh \
		--format yaml --input /seed.yaml

.PHONY: dump
dump:
	docker exec acp ./authorization export \
		--sql.url 'postgres://root@crdb:26257/defaultdb?sslmode=disable' \
		--secret.key KNEcLGdDqpwrXDubqPgDSUkMMsLPXaHh \
		--format yaml > data/seed.yaml
