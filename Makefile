DOCKER_COMPOSE = docker compose

help: ## Outputs this help screen
	@grep -E '(^[a-zA-Z0-9_-]+:.*?##.*$$)|(^##)' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}{printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}' | sed -e 's/\[32m##/[33m/'

build: ## build and start the docker-compose setup for local development
	$(DOCKER_COMPOSE) up --build --force-recreate

up: ## start up the docker-compose setup for local development
	$(DOCKER_COMPOSE) up -d

down: ## start up the docker-compose setup for local development
	$(DOCKER_COMPOSE) down

swagger: ## create swagger api doc
	swag init --parseDependency --parseInternal
	cp docs/swagger.json openapi/openapi.json

lint: ## Lints all code with golangci-lint
	golangci-lint run

lint-fix: ## Lints all code with golangci-lint
	golangci-lint run --fix

# .PHONY declarations https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: up, build, swagger


