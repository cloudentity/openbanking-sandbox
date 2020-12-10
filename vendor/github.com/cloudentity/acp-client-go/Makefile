swagger = docker run --rm -it -e GOPATH=/go \
			-u $(shell id -u ${USER}):$(shell id -g ${USER}) \
			-v $(shell pwd):/go/src \
			-w $(shell pwd)/src quay.io/goswagger/swagger

generate:
	rm -rf client models
	sed -i 's/flow: application/flow: accessCode/g' swagger.yaml
	sed -i 's/flow: password/flow: accessCode/g' swagger.yaml
	${swagger} generate client -f /go/src/swagger.yaml -A acp -t /go/src -q
