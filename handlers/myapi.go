package handlers

import (
	"fmt"
	"github.com/brettman/asyncweb/json"
	"net/http"
)

type Input struct {
	FirstName string
	LastName  string
}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	var input Input
	json.DecodeRequest(r, &input)
	fmt.Fprintf(w, "hello world")
}
