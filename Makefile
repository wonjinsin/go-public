PACKAGE = gorilla

vendor-package:
	go mod vendor

build:
	go build -o bin/gorilla ./api/main.go