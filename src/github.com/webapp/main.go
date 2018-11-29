package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/webapp/controller"
)

func main() {
	templates := populateTemplates()
	controller.Startup(templates)
	// for gzip compressor middleware
	// http.ListenAndServe(":3000", new(middleware.GzipMiddleware))
	http.ListenAndServe(":3000", nil)

}

// populateTemplates for template and subtemplates approaches
// func populateTemplates() *template.Template {
//     result := template.New("templates")
//     const basePath = "templates"
//     template.Must(result.ParseGlob(basePath + "/*.html"))
//     return result
// }

// populateTemplates for template composition approach
func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	layout := template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template blocks directory: " + err.Error())
	}
	fis, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read contents of content directory: " + err.Error())
	}
	for _, fi := range fis {
		f, err := os.Open(basePath + "/content/" + fi.Name())
		if err != nil {
			panic("Failed to open template '" + fi.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read content from file '" + fi.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse contents of '" + fi.Name() + "' as template")
		}
		result[fi.Name()] = tmpl
	}
	return result
}
