package httpserver

import (
	"fmt"
	"net/http"
)

func HttpHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, `{"message": "helloworld""}`)
	})
	return mux
}