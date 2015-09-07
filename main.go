package main

import (
	"fmt"
	"github.com/mongolar/service"
	"net/http"
)

func main() {
	my_service := service.Service{
		Title:   "example_service",
		Version: "0.1",
		Type:    "http",
		Handler: ExampleServe,
	}
	my_service.Serve()
}

func ExampleServe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
	return
}
