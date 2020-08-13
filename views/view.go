package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	//LayoutDir is
	LayoutDir string = "views/layouts/"
	//TemplateExt is
	TemplateExt string = ".gohtml"
)

//NewView is
func NewView(layout string, files ...string) *View {
	filesofLayout := layoutFiles()

	files = append(files,
		filesofLayout...,
	)
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

//Render is
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	err := v.Template.ExecuteTemplate(w, v.Layout, data)
	return err

}

//layoutFiles returns a slice of strings representing
//the layout files used in our applications.
func layoutFiles() []string {
	files, err := filepath.Glob(LayoutDir + "*" + TemplateExt)
	if err != nil {
		panic(err)
	}
	return files

}
