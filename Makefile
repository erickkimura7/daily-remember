BINARY_NAME=server.out

build:
	go mod tidy
	go build -o ${BINARY_NAME}

run:
	go mod tidy
	go build -o ${BINARY_NAME} main.go
	./${BINARY_NAME}
 
clean:
	go clean
	rm ${BINARY_NAME}