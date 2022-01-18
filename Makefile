#>> Docker-compose manage configurations

docker-run:
	@echo "> Docker Run"
	@-docker-compose -f docker-compose.yml up -d


docker-stop:
	@echo "> Docker Stop"
	@docker-compose -f docker-compose.yml stop

docker-build:
	@echo "> Docker Build"
	@-docker-compose -f docker-compose.yml up --build --no-start

docker-remove:
	@echo "> Docker remove stopped containers"
	@-docker-compose -f docker-compose.yml rm --stop --force -v

docker-reset-cache:
	@echo "> Docker remove cached storage"
	@-rm -rf .docker/postgress/tmp
	@-mkdir ".docker/postgress/tmp"

docker-install: docker-stop docker-build docker-run

docker-restart: docker-stop docker-run

docker-clear: docker-stop docker-remove docker-reset-cache
#<< Docker-compose manage configurations


#>> Go Lang Manage
# >> TODO
go-get:
	@echo "> Go get"
	@-go get -d