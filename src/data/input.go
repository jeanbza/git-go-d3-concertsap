package data

import (
    "net/http"
    "html/template"
    "git-go-d3-concertsap/src/database"
    "git-go-d3-concertsap/src/common"
    "log"
)

func GetConcertForm(rw http.ResponseWriter, req *http.Request) {
    type State struct {
        Id   int
        Name string
    }

    type Page struct {
        Title  string
        States []State
    }

    states := []State{}

    var (
        id int
        name string
    )

    rows := database.Select("SELECT id, name FROM state", "concertsap")

    for rows.Next() {
        err := rows.Scan(&id, &name)
        common.CheckError(err)

        states = append(states, State{Id: id, Name: name})
    }

    p := Page{
        Title: "data",
        States: states,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["data.html"] = template.Must(template.ParseFiles("static/html/data/concertform.html", "static/html/data/data.html", "static/html/index.html"))
    tmpl["data.html"].ExecuteTemplate(rw, "base", p)
}

func GetTicketForm(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }

    p := Page{
        Title: "data",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["data.html"] = template.Must(template.ParseFiles("static/html/data/ticketform.html", "static/html/data/data.html", "static/html/index.html"))
    tmpl["data.html"].ExecuteTemplate(rw, "base", p)
}

func Get404(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }

    p := Page{
        Title: "data",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["data.html"] = template.Must(template.ParseFiles("static/html/data/404.html", "static/html/data/data.html", "static/html/index.html"))
    tmpl["data.html"].ExecuteTemplate(rw, "base", p)
}