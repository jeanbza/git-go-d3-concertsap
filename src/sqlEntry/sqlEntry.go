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
    id int
    name string
}

type Retailer struct {
    id int
    name string
}

func GetPage(rw http.ResponseWriter, req *http.Request) {
    // connect to db and query concerts and retailers
    var (
        id int
        name string
    )

    rows := database.Select("SELECT id,name FROM concert", "concertsap")

    for rows.Next() {
        err := rows.Scan(&id, &name)
        common.CheckError(err)
        log.Println(id, name)
    }

    p := Page{
        Title: "sqlEntry",
        //Concerts: ,
        //Retailers: ,
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