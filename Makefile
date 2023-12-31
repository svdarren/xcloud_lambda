.PHONY: build-aws build-azure

all: build-aws build-azure

build-aws:
	CLOUD_TARGET=aws sam build

build-azure:
	ARTIFACTS_DIR=`pwd`/.azure-func \
	CLOUD_TARGET=azure \
	make -C hello-world azure

invoke-aws: build-aws
	sam local invoke -e events/event-n.json

test-bare: build-azure
	CLOUD_PROVIDER=Bash FUNCTIONS_CUSTOMHANDLER_PORT=8088 ./.azure-func/bootstrap
test-aws: build-aws
	sam local start-api --warm-containers EAGER

test-azure: build-azure
	func start --port 8081

test-docker:
	docker compose up --build

hello-all:
	curl \
		localhost:3000/hello \
		localhost:8081/api/hello-world \
		localhost:8080 \
		localhost:8088