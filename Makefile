PACKAGE = gorilla

vendor:
	go mod vendor

build:
	go build -o bin/gorilla ./api/main.go