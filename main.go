package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hello, Docker.\n")
}

func main() {

	http.HandleFunc("/", hello)

	http.ListenAndServe(":3000", nil)
}
