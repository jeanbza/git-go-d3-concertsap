package sqlEntry

import (
    "net/http"
    "html/template"
    "log"
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

func Save(rw http.ResponseWriter, req *http.Request) {
    log.Println(req)

    concertID := req.FormValue("concertID")
    retailerID := req.FormValue("retailerID")
    price := req.FormValue("price")
    timestamp := req.FormValue("timestamp")

    log.Println(concertID, retailerID, price, timestamp)

    http.Redirect(rw, req, "/sqlEntry", http.StatusFound)
}