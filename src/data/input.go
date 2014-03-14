package data

import (
    "net/http"
    "html/template"
)

type Page struct {
    Title string
}

func GetPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "data",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["data.html"] = template.Must(template.ParseFiles("static/html/data/data.html", "static/html/index.html"))
    tmpl["data.html"].ExecuteTemplate(rw, "base", p)
}