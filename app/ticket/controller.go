package ticket

import (
    "net/http"
    "html/template"

    "git-go-d3-concertsap/app/common"

    "github.com/gorilla/mux"
)

func Route(s *mux.Router) {
    s.HandleFunc("/", addHandler)
    s.HandleFunc("/add{_:/?}", addHandler)
}

func addHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "ticket",
        Title:      "Add Controller",
    }

    common.Templates = template.Must(template.ParseFiles("templates/ticket/add.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}