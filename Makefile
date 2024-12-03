.PHONY: up-dev
up-dev:
	docker compose -f deployments/docker-compose.dev.yml up -d --build
up-prod:
	docker compose -f deployments/docker-compose.prod.yml up -d --build --remove-orphans
swag:
	swag init -g cmd/main.go --requiredByDefault 
recreate-dev:
	docker compose -f deployments/docker-compose.dev.yml down && sudo rm -rf deployments/data && docker compose -f deployments/docker-compose.dev.yml up -d