package sqlEntry

import (
    "git-go-d3-concertsap/src/common"
    "git-go-d3-concertsap/src/database"
    "net/http"
    "html/template"
    "log"
)

type Page struct {
    Title string
    Concerts []Concert
    Retailers []Retailer
}

type Concert struct {
    Id int
    Name string
}

type Retailer struct {
    Id int
    Name string
}

func GetPage(rw http.ResponseWriter, req *http.Request) {
    // define concert and retailers arrays to pass into
    // Page instantiation
    concerts := []Concert{}
    retailers := []Retailer{}

    var (
        id int
        name string
    )

    // connect to db and query concerts
    rows := database.Select("SELECT id,name FROM concert", "concertsap")

    for rows.Next() {
        err := rows.Scan(&id, &name)
        common.CheckError(err)
        log.Println(id, name)
        concerts = append(concerts, Concert{ Id:id, Name:name })
    }

    // connect to db and query retailers
    rows = database.Select("SELECT id,name FROM retailer", "concertsap")

    for rows.Next() {
        err := rows.Scan(&id, &name)
        common.CheckError(err)
        log.Println(id, name)
        retailers = append(retailers, Retailer{ Id:id, Name:name })
    }
    p := Page{
        Title: "sqlEntry",
        Concerts: concerts,
        Retailers: retailers,
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