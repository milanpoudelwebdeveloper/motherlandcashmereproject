package controllers

import (
	"../views"
)

//NewStatic is
func NewStatic() *Static {
	return &Static{
		Home:    views.NewView("bootstrap", "static/home"),
		Contact: views.NewView("bootstrap", "static/contact"),
	}
}

//Static is
type Static struct {
	Home    *views.View
	Contact *views.View
}
