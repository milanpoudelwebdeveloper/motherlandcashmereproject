package controllers

import (
	"fmt"
	"net/http"

	"../models"
	"../views"
)

//NewUsers is used to create a new users controller.
//This function will panic if templates are not parsed correctly
//and should only be used during the initial setup.

//NewUsers is
func NewUsers(us *models.UserService) *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "users/new"),
		us:      us,
	}
}

//Users is
type Users struct {
	NewView *views.View
	us      *models.UserService
}

// //GET/signup

//New is used to render the form where the user can create a new user account
func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

//SignupForm is
type SignupForm struct {
	Name     string `schema:"name"`
	Email    string `schema:"email"`
	Password string `schema:"password"`
}

//POST/signup

//Create is used to create new user account
//Create is used to process the signup form where users submits it
func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	var form SignupForm
	if err := parseForm(r, &form); err != nil {
		panic(err)
	}
	user := models.User{
		Name:     form.Name,
		Email:    form.Email,
		Password: form.Password,
	}
	if err := u.us.Create(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, user)

}
