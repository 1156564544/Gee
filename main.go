package main

import (
	"Gee/gee"
	"fmt"
	"net/http"
)

func main() {
	r := gee.New()

	r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "URL.Path = %q\n", request.URL.Path)
	})
	r.Get("/hello", func(writer http.ResponseWriter, request *http.Request) {
		for k, v := range request.Header {
			fmt.Fprintf(writer, "Header[%q] = %q\n", k, v)
		}
	})

	r.Run(":9999")
}
