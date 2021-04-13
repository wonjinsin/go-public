PACKAGE = chatapp

vendor:
	go mod vendor

build:
	go build -o bin/chatapp ./api/main.go