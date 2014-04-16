package api

import (
    "net/http"
    "fmt"
    "encoding/json"

    "git-go-d3-concertsap/app/common"
    "git-go-d3-concertsap/app/band"

    "github.com/gorilla/mux"
)

func ConcertBandHandler(rw http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id := params["id"]

    json, err := json.Marshal(band.FindBandConcert(id))
    common.CheckError(err)

    fmt.Fprintf(rw, "%v", string(json))
}