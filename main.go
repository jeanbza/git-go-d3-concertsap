package main

import (
    "git-go-d3-concertsap/src/home"
    "git-go-d3-concertsap/src/common"
    "net/http"
)

func main() {
    http.HandleFunc("/", homePage)

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    common.CheckError(err)
}

func homePage(rw http.ResponseWriter, req *http.Request) {
    home.GetPage(rw, req)
}