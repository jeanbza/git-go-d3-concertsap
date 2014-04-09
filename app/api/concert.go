package api

import (
    "net/http"
    "fmt"
    "encoding/json"

    "git-go-d3-concertsap/app/common"
    "git-go-d3-concertsap/app/band"
)

func ConcertBandHandler(rw http.ResponseWriter, req *http.Request) {
    json, err := json.Marshal(band.FindBandConcert())
    common.CheckError(err)

    fmt.Fprintf(rw, "%v", string(json))
}