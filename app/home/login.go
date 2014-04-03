package home

import (
    "net/http"
    "html/template"
    "log"

    "git-go-d3-concertsap/app/common"

    "github.com/gorilla/mux"
)

func Route(s *mux.Router) {
    s.HandleFunc("/", loginHandler)
    s.HandleFunc("/validate{_:/?}", validateHandler).Methods("POST")
}

func loginHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "login",
        Title:      "Login",
    }

    common.Templates = template.Must(template.ParseFiles("templates/home/login.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func validateHandler(rw http.ResponseWriter, req *http.Request) {

    log.Println("validate!")

    http.Redirect(rw, req, "/login/", http.StatusFound)
}