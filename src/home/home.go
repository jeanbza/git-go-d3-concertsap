package home

import (
    "net/http"
    "html/template"
)

type Page struct {
    PageName    string
    Title       string
}

func GetPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        PageName:   "home",
        Title:      "Home",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["home.html"] = template.Must(template.ParseFiles("templates/home.html", "templates/index.html"))
    tmpl["home.html"].ExecuteTemplate(rw, "base", p)
}