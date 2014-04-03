package band

import (
    "net/http"
    "html/template"
    "regexp"
    "strconv"
    "strings"

    "git-go-d3-concertsap/app/common"
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/admin/concert"

    "github.com/gorilla/mux"
)

func Route(s *mux.Router) {
    s.HandleFunc("/", viewAllHandler)
    s.HandleFunc("/{id:[0-9]+}{_:/?}", viewOneHandler)
    s.HandleFunc("/view/{id:[0-9]+}", viewOneHandler)
}

func viewAllHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
        Bands       []Band
    }

    bands := FindAll()

    p := Page{
        PageName:   "band",
        Title:      "View All Bands",
        Bands:      bands,
    }

    common.Templates = template.Must(template.ParseFiles("templates/admin/band/viewAll.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func viewOneHandler(rw http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id := params["id"]

    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "band",
        Title:      "View One Controller: "+id,
    }

    common.Templates = template.Must(template.ParseFiles("templates/admin/band/viewOne.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}