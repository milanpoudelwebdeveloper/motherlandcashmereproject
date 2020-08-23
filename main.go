package main

import (
	"fmt"
	"net/http"

	"./controllers"
	"./models"
	"github.com/gorilla/mux"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "milanpoudel"
	dbname   = "lenslocked_dev"
)

func must(err error) {
	if err != nil {
		panic(err)
	}
}
func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	us, err := models.NewUserService(psqlInfo)
	must(err)
	defer us.Close()
	us.AutoMigrate()
	staticC := controllers.NewStatic()
	usersC := controllers.NewUsers(us)
	r := mux.NewRouter()
	r.Handle("/", staticC.Home).Methods("GET")
	r.Handle("/contact", staticC.Contact).Methods("GET")
	r.HandleFunc("/signup", usersC.New).Methods("GET")
	r.HandleFunc("/signup", usersC.Create).Methods("POST")
	http.ListenAndServe(":8054", r)
}
