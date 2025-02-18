package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", homepage)
	http.ListenAndServe("localhost:1000", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Implementing a demo Api server")
	fmt.Println("Checking If we hit the right end point for Homepage")

}
