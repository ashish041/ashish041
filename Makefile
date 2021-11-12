# make all to run the app in local
all: go-install up

go-install:
	sh install-go.sh

up: build run test callApi

build:
	sh ./build.sh
run:
	sh ./run.sh
test:
	sh ./test.sh
callApi:
	sh ./callApi.sh

# make docker-all to run the app in docker
docker-all: docker-build docker-run callApi
docker-build:
	sh ./docker-build.sh
docker-run:
	sh ./docker-run.sh
