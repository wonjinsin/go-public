PACKAGE = gorilla
OS = ${GOOS}


vendor-package:
	go mod vendor

build:
	GOOS=$(OS) go build -o bin/gorilla ./api/main.go
