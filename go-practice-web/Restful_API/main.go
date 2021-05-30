package main

import (
	"net/http"

	"github.com/wonjinsin/go-public/go-practice-web/Restful_API/myapp"
)

func main() http.Handler {
	http.ListenAndServe(":3000", myapp.NewHandler())
}
