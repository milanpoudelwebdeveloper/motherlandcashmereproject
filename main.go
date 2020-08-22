package main

import (
	"net/http"

	"./controllers"

	"github.com/gorilla/mux"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers()
	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":8054", r)
}
