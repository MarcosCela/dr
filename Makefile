## Makefile directives
.PHONY: help build test populate registry
.DEFAULT_GOAL := help
help: ## Shows this help message and all the other targets with nice colored output
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build: test ## Removes the go binary, and then builds another go binary
	@rm -f dr
	@go build

registry: ## Launch a local docker registry for testing
	@docker-compose -f docker-compose.yaml up -d

populate: ## Downloads busybox image to the testing repository
	@docker pull busybox:1.29
	@docker pull alpine:latest
	@docker tag busybox:1.29 localhost:5000/busybox:1.29
	@docker tag busybox:1.29 localhost:5000/busybox:mytag
	@docker tag busybox:1.29 localhost:5000/busybox/testing:latest
	@docker tag busybox:1.29 localhost:5000/busybox/testing:mytag
	@docker tag busybox:1.29 localhost:5000/busybox/another/subpath/testing:anothertag
	@docker tag busybox:1.29 localhost:5000/another/busy/subpath:mytag
	@docker tag alpine:latest localhost:5000/alpine:latest
	@docker push localhost:5000/busybox/busy/subpath:mytag
	@docker push localhost:5000/busybox:1.29
	@docker push localhost:5000/busybox/testing:latest
	@docker push localhost:5000/busybox/testing:mytag
	@docker push localhost:5000/busybox/another/subpath/testing:anothertag
	@docker push localhost:5000/alpine:latest

test: format ## Test using the default 'go test' command
	@go test ./... -v

format: ## Format the code using go fmt
	@go fmt ./...

clean: ## Cleans the docker compose testing images
	@rm -rf ./registry

install: build ## Copy the executable to ${HOME}/bin/dr
	@cp dr $(HOME)/bin/dr
	@chmod +x $(HOME)/bin/dr
