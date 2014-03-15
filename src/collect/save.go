package collect

import (
     "net/http"
     "log"
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