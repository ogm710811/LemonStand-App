package controller

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/webapp/viewmodel"
)

type home struct {
	homeTemplate         *template.Template
	loginTemplate        *template.Template
	standLocatorTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()

	//check for method in the form
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("Error logging in: %v", err))
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		if email == "ogm@gmail.com" && password == "12345" {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		}

		vm.Email = email
		vm.Password = password
	}
	w.Header().Add("Content-Type", "text/html")
	h.loginTemplate.Execute(w, vm)
}
