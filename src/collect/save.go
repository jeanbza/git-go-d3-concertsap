package collect

import (
     "net/http"
     "log"
     "git-go-d3-concertsap/src/database"
     "git-go-d3-concertsap/src/common"
)

func SaveForm(rw http.ResponseWriter, req *http.Request) {
    log.Println(req)

    concertID := req.FormValue("concertID")
    retailerID := req.FormValue("retailerID")
    price := req.FormValue("price")
    timestamp := req.FormValue("timestamp")

    log.Println(concertID, retailerID, price, timestamp)

    http.Redirect(rw, req, "/collect/ticket", http.StatusFound)
}