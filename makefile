BINARY_NAME=tooter

build:
	GOARCH=amd64 GOOS=linux go build -o bin/${BINARY_NAME} main.go

run: build
		go run bin/${BINARY_NAME}

clean:
	go clean
	rm ./bin/${BINARY_NAME}