package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "git-go-d3-concertsap/app/home"
    "git-go-d3-concertsap/app/concert"
    "git-go-d3-concertsap/app/database"
)

func main() {
    db.InitDb()

    r := mux.NewRouter()

    http.Handle("/", r)
    
    r.HandleFunc("/", HandleHome).Methods("GET")

    // Concert
    s := r.PathPrefix("/concert").Subrouter()
    concert.Route(s)

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    http.ListenAndServe(":8080", nil)
}

func HandleHome(rw http.ResponseWriter, req *http.Request) {
    home.GetPage(rw, req)
}