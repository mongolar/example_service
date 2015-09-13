package main

import (
	"fmt"
	"github.com/mongolar/service"
	"net/http"
)

func main() {
	my_service := service.New()
	my_service.Handler = http.HandlerFunc(ExampleServe)
	my_service.Serve()
}

func ExampleServe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
	return
}
