package main

import (
	"github.com/brettman/RestServer"
	"github.com/brettman/asyncweb/handlers"
	"log"
	"net/http"
)

var routes = []server.Route{
	{"Hello", "POST", "/hello", handlers.HelloWorld},
}

func main() {
	router := server.NewRouter(routes)
	log.Fatal(http.ListenAndServe(":8080", router))
}
