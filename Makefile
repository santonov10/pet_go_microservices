docker_compose_up_recreate:
	docker-compose up --build --force-recreate --no-deps

docker_compose_up:
	docker-compose up