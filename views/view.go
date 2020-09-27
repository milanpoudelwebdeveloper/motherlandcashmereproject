package views

import (
	"html/template"
	"net/http"
	"path/filepath"
)

var (
	//TemplateDir is
	TemplateDir string = "views/"
	//LayoutDir is
	LayoutDir string = "views/layouts/"
	//TemplateExt is
	TemplateExt string = ".gohtml"
)

//NewView is
func NewView(layout string, files ...string) *View {
	addTemplatePath(files)
	addTemplateExt(files)
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

func (v *View) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := v.Render(w, nil); err != nil {
		panic(err)
	}
}

//Render is
func (v *View) Render(w http.ResponseWriter, data interface{}) error {
	w.Header().Set("Content-Type", "text/html charset=utf-8")
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

//addTemplatePath takes in a slice of strings
//representing the file paths for templates and it prepends
//the TemplateDir directory to each string in the slice

//Eg.the input {"home"} would result in the output
//{"views/home"}if templateDir=="views/"

func addTemplatePath(files []string) {
	for i, f := range files {
		files[i] = TemplateDir + f
	}

}

//addTemplateExt takes in a slice of strings
//representing the file paths for the templates and it appends
//the TemplateExt extenstion to each string in the slice

//Eg. the input {"home"} would result in the output
//{home.gohtml } if templateExt==".gohtml"
//
func addTemplateExt(files []string) {
	for i, f := range files {
		files[i] = f + TemplateExt
	}
}
