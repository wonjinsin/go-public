PACKAGE = giraffe
OS = ${GOOS}


vendor-package:
	go mod vendor

build:
	GOOS=$(OS) go build -o bin/giraffe ./api/main.go && \
	cp -fpR key bin/.
