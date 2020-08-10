package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/dogs/", handlerFunc)
	http.ListenAndServe(":8080", nil)

}

func handlerFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Welcome to our copy site</h1>")
}
