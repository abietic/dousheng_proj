build:
	sh scripts/total_build.sh
up: build
	docker-compose up -d
down:
	docker-compose down --rmi local
rup:
	echo "up but don't rebuild"
	docker-compose up -d
help:
	@echo "'make build' to rebuild all go modules"
	@echo "'make up' to rebuid and up docker compose"
	@echo "'make rup' to up docker compose only"
	@echo "'make down' ro down docker compose"