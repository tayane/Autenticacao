package api

import (
	"autenticacao/services"
	"html/template"
	"net/http"
	"path/filepath"
)

var templates = template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if err := services.CreateUser(username, password); err != nil {
			http.Error(w, "Unable to create user", http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	templates.ExecuteTemplate(w, "register.html", nil)
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		username := r.FormValue("username")
		password := r.FormValue("password")

		if !services.AuthenticateUser(username, password) {
			http.Error(w, "Invalid username or password", http.StatusUnauthorized)
			return
		}

		templates.ExecuteTemplate(w, "welcome.html", nil)
		return
	}

	templates.ExecuteTemplate(w, "login.html", nil)
}
