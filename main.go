package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("hello world")

	// define a handler function
	// the first argument is the URL, '/'
	// the second is h1, a handler defined above the function

	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "hello world\n")
		io.WriteString(w, r.Method)
	}
	http.HandleFunc("/", h1)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
