BINARY_NAME=server.out
DIRECTORY_BIN_NAME=
MOCK_GEN_PATH=$(HOME)/go/bin/mockgen

build:
	go mod tidy
	go build -o ${BINARY_NAME}

build-prod:
	go mod tidy
	go build -o ${BINARY_NAME}

run:
	go mod tidy
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

gen-mocks:
	$(MOCK_GEN_PATH) -destination=./mocks/notificationevent/service.go -source=notificationEvent/service.go
	$(MOCK_GEN_PATH) -destination=./mocks/notificationevent/model.go -source=notificationEvent/model.go

test:
	go test ./tests/... -v -coverpkg=./...