package data

import (
    "net/http"
    "html/template"
)

type Page struct {
    Title string
}

func GetConcertForm(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "data",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["data.html"] = template.Must(template.ParseFiles("static/html/data/concertform.html", "static/html/data/data.html", "static/html/index.html"))
    tmpl["data.html"].ExecuteTemplate(rw, "base", p)
}

func GetTicketForm(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "data",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["data.html"] = template.Must(template.ParseFiles("static/html/data/ticketform.html", "static/html/data/data.html", "static/html/index.html"))
    tmpl["data.html"].ExecuteTemplate(rw, "base", p)
}

func Get404(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "data",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["data.html"] = template.Must(template.ParseFiles("static/html/data/404.html", "static/html/data/data.html", "static/html/index.html"))
    tmpl["data.html"].ExecuteTemplate(rw, "base", p)
}