.PHONY: build
build:
	make bank-deps
	docker-compose build

.PHONY: bank-deps
bank-deps:
	cd bank && go mod vendor

.PHONY: run
run:
	docker-compose up -d
	./scripts/wait.sh
	make seed # todo add import to acp start cmd

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
