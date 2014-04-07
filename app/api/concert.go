package api

import (
    "net/http"
    "fmt"
)

func ConcertBandHandler(rw http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(rw, "boom")
}