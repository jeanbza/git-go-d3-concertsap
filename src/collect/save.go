package collect

import (
    // "git-go-d3-concertsap/src/common"
    // "git-go-d3-concertsap/src/database"
     "net/http"
    // "html/template"
     "log"
)

func Save(rw http.ResponseWriter, req *http.Request) {
    log.Println(req)

    concertID := req.FormValue("concertID")
    retailerID := req.FormValue("retailerID")
    price := req.FormValue("price")
    timestamp := req.FormValue("timestamp")

    log.Println(concertID, retailerID, price, timestamp)

    http.Redirect(rw, req, "/sqlEntry", http.StatusFound)
}