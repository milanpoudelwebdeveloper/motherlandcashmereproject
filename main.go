package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	fmt.Fprint(w, `<h1>Welcome to the awesome site</h1>`)
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	fmt.Fprint(w, `To get in touch, please send an email to <a href="mailto:support@lenslocked.com">support@lenslocked.com</a>`)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", r)

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi, Welcome here")
}
