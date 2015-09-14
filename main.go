package main

import (
	"fmt"
	"github.com/mongolar/service"
	"net/http"
)

func main() {
	service.Handler(http.HandlerFunc(ExampleServe))
	service.Serve()
}

func ExampleServe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
	return
}
