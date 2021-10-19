.PHONY: lint develop
.DEFAULT_GOAL := help

GOPATH = $(shell go env GOPATH)

lint: ## Lint the source code
	revive ./...

gen-proto: ## Generate the protobuf files
	protoc --proto_path=${GOPATH}/src:. --twirp_out=. --go_out=. internal/apps/postservice/rpc/post.proto

develop: ## Run the containers
	docker-compose up -d

develop-it: ## Run the containers in interactive mode
	docker-compose up

develop-clean: ## Run the containers with --build and --remove-orphans flag
	docker-compose up -d --build --remove-orphans

develop-clean-it: ## Run the containers with --build and --remove-orphans flag in interactive mode
	docker-compose up --build --remove-orphans

down: ## Stop the containers
	docker-compose down

stop: ## Stop the containers and remove the containers
	docker-compose down

post-service-database-shell: ## Run the post service database shell
	docker-compose exec post-service-database psql -U postgres

post-service-redis-shell: ## Run the post service database shell
	docker-compose exec post-service-redis redis-cli

# Source: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Displays all the available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)