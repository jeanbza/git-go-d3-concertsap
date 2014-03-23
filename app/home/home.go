package home

import (
    "net/http"
    "html/template"

    "git-go-d3-concertsap/app/common"
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

    common.Templates = template.Must(template.ParseFiles("templates/home/home.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}