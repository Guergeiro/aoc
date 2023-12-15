.PHONY: test

DOCKER_HOME=/home/app

all:
	docker run \
		--interactive --tty --rm --volume ${PWD}:${DOCKER_HOME} \
		--workdir ${DOCKER_HOME} \
		golang:1.21 go run main.go

build:
	docker run \
		--interactive --tty --rm --volume ${PWD}:${DOCKER_HOME} \
		--workdir ${DOCKER_HOME} \
		golang:1.21 go build

test:
	docker run \
		--interactive --tty --rm --volume ${PWD}:${DOCKER_HOME} \
		--workdir ${DOCKER_HOME} \
		golang:1.21 go test ./test/**

run:
	docker run \
		--interactive --tty --rm --volume ${PWD}:${DOCKER_HOME} \
		--workdir ${DOCKER_HOME} \
		golang:1.21 go run main.go $(filter-out $@,$(MAKECMDGOALS))

%:
	@:
