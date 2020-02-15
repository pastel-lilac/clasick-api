DOCKER_COMPOSE_DEV_PATH=./deployments/docker-compose.dev.yaml
prepare-dev:
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) up -d core_db
up-core-dev:
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) build
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) up -d core
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) logs -f core
log-core-dev:
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) logs -f core
ps-dev:
	sudo docker-compose -f $(DOCKER_COMPOSE_DEV_PATH) ps
