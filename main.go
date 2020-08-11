package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gorilla/mux"
)

var homeTemplate *template.Template

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

var contactTemplate *template.Template

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	if err := contactTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}

func main() {

	var err error
	homeTemplate, err = template.ParseFiles("views/home.gohtml")
	if err != nil {
		panic(err)
	}

	contactTemplate, err = template.ParseFiles("views/contact.gohtml")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", r)

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi, Welcome here")
}
