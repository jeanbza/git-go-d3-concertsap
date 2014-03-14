package sqlEntry

import (
    "net/http"
    "html/template"
)

type Page struct {
    Title string
}

func GetPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "sqlEntry",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["sqlEntry.html"] = template.Must(template.ParseFiles("static/html/sqlEntry.html", "static/html/index.html"))
    tmpl["sqlEntry.html"].ExecuteTemplate(rw, "base", p)
}