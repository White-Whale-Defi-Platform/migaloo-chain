all: migalood-upgrade-env

migalood-upgrade-env:
	docker build --platform linux/amd64 --no-cache --build-arg UID=$(shell id -u) --build-arg GID=$(shell id -g) --build-arg entrypoint=cosmovisor-entrypoint.sh --tag migaloo/migalood-upgrade-env .

.PHONY: all migalood-upgrade-env
