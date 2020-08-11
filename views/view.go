package views

import "html/template"

//NewView is
func NewView(files ...string) *View {
	files = append(files, "views/layouts/footer.gohtml")
	t, err := template.ParseFiles(files...)
	if err != nil {
		panic(err)

	}
	return &View{
		Template: t,
	}
}

//View is
type View struct {
	Template *template.Template
}
