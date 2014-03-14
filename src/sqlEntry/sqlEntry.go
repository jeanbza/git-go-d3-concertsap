package sqlEntry

import (
    "net/http"
    "html/template"
    "fmt"
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

func submitHandler(rw http.ResponseWriter, req *http.Request) {
    concertID := req.FormValue("concertID")
    retailerID := req.FormValue("retailerID")
    price := req.FormValue("price")
    timestamp := req.FormValue("timestamp")

    fmt.Printf(concertID)
    fmt.Printf(retailerID)
    fmt.Printf(price)
    fmt.Printf(timestamp)

    http.Redirect(rw, req, "/sqlEntry", http.StatusFound)
}