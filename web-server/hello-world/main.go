package main

import (
	"fmt"
	"net/http"
)

const PORT = ":8080"

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(PORT, nil)
}

func greet(w http.ResponseWriter, r *http.Request) {
	msg := "hello world"
	fmt.Fprint(w, msg)
}
