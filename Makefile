.PHONY: up-dev
up-dev:
	docker compose -f deployments/docker-compose.dev.yml up -d
swag:
	swag init -g cmd/main.go --requiredByDefault 