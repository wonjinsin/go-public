package main

import (
	"net/http"

	"github.com/wonjinsin/go-practice-web/Restful_API/app"
)

func main() http.Handler {
	http.ListenAndServe(":3000", app.NewHandler())
}
