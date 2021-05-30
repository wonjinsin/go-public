package main

import (
	"net/http"

	"github.com/wonjinsin/go-public/go-practice-web/test/myapp"
)

func main() {

	http.ListenAndServe(":3000", myapp.NewHttpHandler())
}
