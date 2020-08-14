package controllers

import (
	"../views"
)

//NewStatic is
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "views/static/home.gohtml"),
		Contact: views.NewView("bootstrap", "views/static/contact.gohtml"),
	}
}

//Static is
type Static struct {
	Home    *views.View
	Contact *views.View
}
