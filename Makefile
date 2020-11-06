.PHONY: build
build:
	docker-compose build

.PHONY: download-deps
download-deps:
	go mod vendor

.PHONY: run
run:
	docker-compose up -d
	./scripts/wait.sh
	make seed # todo add import to acp start cmd

.PHONY: lint
lint:
	golangci-lint run --fix --deadline=5m ./...

.PHONY: clean
clean:
	docker-compose down

.PHONY: seed
seed:
	docker exec acp ./authorization import \
		--sql.url 'postgres://user:password@postgres/authorization?sslmode=disable' \
		--secret.key KNEcLGdDqpwrXDubqPgDSUkMMsLPXaHh \
		--format yaml --input /seed.yaml

.PHONY: dump
dump:
	docker exec acp ./authorization export \
		--sql.url 'postgres://user:password@postgres/authorization?sslmode=disable' \
		--secret.key KNEcLGdDqpwrXDubqPgDSUkMMsLPXaHh \
		--format yaml > data/seed.yaml
