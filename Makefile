.PHONY: build
build:
	docker-compose build --parallel

.PHONY: download-deps
download-deps:
	go mod vendor
	go mod tidy

.PHONY: replace-hosts
replace-hosts:
	./scripts/replace_hosts.sh

.PHONY: run
run: replace-hosts
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

swagger = docker run --rm -it -e GOPATH=/go \
			-u $(shell id -u ${USER}):$(shell id -g ${USER}) \
			-v $(shell pwd):/go/src \
			-w $(shell pwd)/src quay.io/goswagger/swagger

generate:
	rm -rf client models
	${swagger} generate client -f /go/src/swagger.yaml -A openbanking -t /go/src -q
