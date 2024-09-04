# Variables
DOCKER_COMPOSE_FILE := docker-compose.yml

# Default target: build and run the application
all: build up

# Build Docker images
build:
	@echo "Building Docker images..."
	docker compose -f $(DOCKER_COMPOSE_FILE) build

# Start the application in detached mode
up:
	@echo "Starting Docker containers..."
	docker compose -f $(DOCKER_COMPOSE_FILE) up -d

# Stop and remove all containers, networks, and volumes
remove:
	@echo "Stopping and removing Docker containers..."
	docker compose -f $(DOCKER_COMPOSE_FILE) down


clear-images:
	@echo "Destroying images..."
	docker system prune -a

migrate:
	@echo "Running migrations..."
	docker-compose run migrations

# Clear all Docker volumes
clear-volumes:
	@echo "Removing all Docker volumes..."
	docker volume prune -f

destroy: clear-volumes clear-images

# Full cleanup: remove containers and clear volumes
clean: remove clear-volumes

# Build, run, remove containers, and clear volumes (full workflow)
run: all clean

test: 
	@echo "Starting tests..."
	./setup-test-db.sh


.PHONY: all build up remove clear-volumes clean run
