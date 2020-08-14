package controllers

import (
	"net/http"

	"../views"
)

//NewUsers is used to create a new users controller.
//This function will panic if templates are not parsed correctly
//and should only be used during the initial setup.

//NewUsers is
func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.gohtml"),
	}
}

//Users is
type Users struct {
	NewView *views.View
}

//New is
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}
