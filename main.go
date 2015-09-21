package main

import (
	"fmt"
	"github.com/mongolar/microservice/server"
	"net/http"
)

func main() {
	server.Handler(http.HandlerFunc(ExampleServe))
	server.Serve()
}

func ExampleServe(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
	return
}
