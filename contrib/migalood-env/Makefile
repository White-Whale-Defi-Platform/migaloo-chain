all: migalood-upgrade-env

migalood-upgrade-env:
	docker build --platform linux/amd64 --no-cache --build-arg entrypoint=cosmovisor-entrypoint.sh --tag migaloo/migalood-upgrade-env .

.PHONY: all migalood-upgrade-env
