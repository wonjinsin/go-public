PACKAGE = chatapp

vendor:
	@echo "chatapp"

build:
	go build -o bin/chatapp ./api/main.go