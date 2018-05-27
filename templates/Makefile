makefile_dir 	:= $(abspath $(shell pwd))

docker_service	:= bi-gpg
docker_tag		:= $(shell cat docker-compose.yml | grep image | sed 's/.*image://' | tr -d ' ')
docker_src	 	:= /src

bumpversion_c	:= docker run --rm -v $(makefile_dir):/dist bi-gpg:1.0.0

.PHONY: list docker-build docker-run docker-push

list:
	@grep '^[^#[:space:]].*:' Makefile | grep -v ':=' | grep -v '^\.' | sed 's/:.*//g' | sed 's/://g' | sort

docker-tag:
	@echo $(docker_tag)

docker-up:
	docker-compose up $(docker_service)

docker-stop:
	docker-compose stop $(docker_service)

docker-build:
	docker-compose build $(docker_service)

docker-build-nocache:
	docker-compose build --no-cache $(docker_service)

docker-bash:
	docker-compose run --rm $(docker_service) bash

docker-push:
	docker login $(DOCKER_REPO) --username '$(DOCKER_USERNAME)' --password '$(DOCKER_PASSWORD)'
	docker push $(docker_tag)

bumpversion-patch:
	@ $(bumpversion_c) make patch
