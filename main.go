package main

import (
	"net/http"

	"github.com/danielvladco/github-ci-cd-workflow/internal/httpserver"
)

func main() {
	hdl := httpserver.HttpHandler()
	http.ListenAndServe(":8080", hdl)
}
