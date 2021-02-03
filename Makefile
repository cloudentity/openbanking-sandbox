export DOCKER_BUILDKIT=1
export COMPOSE_DOCKER_CLI_BUILD=1

.EXPORT_ALL_VARIABLES: ;

.PHONY: build
build:
	docker-compose -f docker-compose.yaml -f docker-compose.build.yaml build

.PHONY: replace-hosts
replace-hosts:
	./scripts/replace_hosts.sh

.PHONY: run-dev
run-dev: replace-hosts
	docker-compose -f docker-compose.yaml -f docker-compose.build.yaml up -d
	./scripts/wait.sh
	make seed

.PHONY: run
run: replace-hosts
	docker-compose up -d --no-build
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

swagger = docker run --rm -it -e GOPATH=/go \
			-u $(shell id -u ${USER}):$(shell id -g ${USER}) \
			-v $(shell pwd):/go/src \
			-w $(shell pwd)/src quay.io/goswagger/swagger

generate:
	rm -rf openbanking/accountinformation/*
	rm -rf opebanking/paymentinitiation/*
	${swagger} generate client -f /go/src/api/accounts.yaml -A OpenbankingAccountsClient -t /go/src/openbanking/accountinformation 
	${swagger} generate client -f /go/src/api/paymentinitiation.yaml -A OpenbankingPaymentsClient -t /go/src/openbanking/paymentinitiation 
