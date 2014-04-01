package ticket

import (
    "net/http"
    "html/template"
    "time"

    "git-go-d3-concertsap/app/common"
    "git-go-d3-concertsap/app/concert"
    "git-go-d3-concertsap/app/retailer"

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
        Concerts    []concert.Concert
        Retailers   []retailer.Retailer
        Now         string
    }

    concerts := concert.FindAll()
    retailers := retailer.FindAll()
    t := time.Now()

    p := Page{
        PageName:   "ticket",
        Title:      "Add Controller",
        Concerts:   concerts,
        Retailers:  retailers,
        Now:        t.Format("2006-01-02 15:04:05"),
    }

    common.Templates = template.Must(template.ParseFiles("templates/ticket/add.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}