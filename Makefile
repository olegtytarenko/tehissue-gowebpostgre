#>> Docker-compose manage configurations
.PHONY: docker-run
docker-run:
	@echo "> Docker Run"
	@-docker-compose -f docker-compose.yml up

.PHONY: docker-run-demonize
docker-run-demonize:
	@echo "> Docker Run Demonize"
	@-docker-compose -f docker-compose.yml up -d

.PHONY: docker-stop
docker-stop:
	@echo "> Docker Stop"
	@docker-compose -f docker-compose.yml stop

.PHONY: docker-build
docker-build:
	@echo "> Docker Build"
	@-docker-compose -f docker-compose.yml up --build --no-start

.PHONY: docker-remove
docker-remove:
	@echo "> Docker remove stopped containers"
	@-docker-compose -f docker-compose.yml rm --stop --force -v

.PHONY: docker-reset-cache
docker-reset-cache:
	@echo "> Docker remove cached storage"
	@-rm -rf .docker/postgress/tmp
	@-mkdir ".docker/postgress/tmp"

.PHONY: docker-install
docker-install: docker-stop docker-build docker-run-demonize

.PHONY: docker-restart
docker-restart: docker-stop docker-run

.PHONY: docker-clear
docker-clear: docker-stop docker-remove docker-reset-cache
#<< Docker-compose manage configurations


#>> Go Lang Manage
# >> TODO
.PHONY: setup
setup:
	@echo "> Go get"
	@-go get -d -v ./...
	@-go install -v ./...

.PHONY: build
build:
	@echo "> Go build"
	go build -o bin/app cmd/app/main.go

.PHONY: clean
clean:
	go clean


.PHONY: install
setup-run: setup build clean

.PHONY: demo
demo: docker-stop docker-build docker-run