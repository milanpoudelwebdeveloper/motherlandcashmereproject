package views

import "html/template"

//NewView is
func NewView(layout string, files ...string) *View {

	files = append(files, "views/layouts/bootstrap.gohtml", "views/layouts/footer.gohtml", "views/layouts/navbar.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)
	}
	return &View{
		Template: t,
		Layout:   layout,
	}
}

//View is
type View struct {
	Template *template.Template
	Layout   string
}
