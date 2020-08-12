package main

import (
	"fmt"
	"net/http"

	"./views"

	"github.com/gorilla/mux"
)

var (
	homeView    *views.View
	contactView *views.View
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	err := homeView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}

}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
	err := contactView.Template.Execute(w, nil)
	if err != nil {
		panic(err)
	}

}

func main() {
	homeView = views.NewView("views/home.gohtml")
	contactView = views.NewView("views/contact.gohtml")
	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)
	http.ListenAndServe(":8080", r)

}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hi, Welcome here")
}
