.PHONY: build-aws build-azure

all: build-aws build-azure

build-aws:
	CLOUD_TARGET=aws sam build

build-azure:
	ARTIFACTS_DIR=`pwd`/.azure-func \
	CLOUD_TARGET=azure \
	make -C hello-world

invoke-aws: build-aws
	sam local invoke -e events/event-n.json

test-aws: build-aws
	sam local start-api --warm-containers EAGER

test-azure: build-azure
	func start --port 8081

test-docker:
	docker compose up --build

hello-all:
	curl localhost:3000/hello localhost:8081/api/hello-world localhost:8080