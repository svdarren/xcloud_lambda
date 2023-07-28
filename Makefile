.PHONY: build-aws build-azure

build-aws:
	CLOUD_TARGET=aws sam build

build-azure:
	ARTIFACTS_DIR=`pwd`/.azure-func \
	CLOUD_TARGET=azure \
	make -C hello-world
