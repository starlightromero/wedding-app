.DEFAULT_GOAL := start-dev

.PHONY: start-dev
start-dev:
	docker-compose -f docker-compose.dev.yml up

.PHONY: stop-dev
stop-dev:
	docker-compose -f docker-compose.dev.yml down

.PHONY: reload-dev
reload-dev: stop-dev start-dev

.PHONY: hard-reload-dev
hard-reload-dev: stop-dev rmi start-dev

.PHONY: rmi
rmi:
	docker rmi wedding-app
