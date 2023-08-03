build-HelloWorldFunction:
	CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -v -o $(ARTIFACTS_DIR)/bootstrap -tags ${CLOUD_TARGET} .

azure:
	CGO_ENABLED=0 go build -v -o $(ARTIFACTS_DIR)/bootstrap -tags ${CLOUD_TARGET} .