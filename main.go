package main

import (
    "net/http"

    "git-go-d3-concertsap/app/home"
    "git-go-d3-concertsap/app/concert"
    "git-go-d3-concertsap/app/state"
    "git-go-d3-concertsap/app/band"
    "git-go-d3-concertsap/app/ticket"
    "git-go-d3-concertsap/app/retailer"
    "git-go-d3-concertsap/app/database"

    "github.com/gorilla/mux"
)

func main() {
    db.InitDb()

    r := mux.NewRouter()

    http.Handle("/", r)
    
    r.HandleFunc("/", HandleHome).Methods("GET")

    // Concert
    s := r.PathPrefix("/concert").Subrouter()
    concert.Route(s)

    // Retailer
    s = r.PathPrefix("/retailer").Subrouter()
    retailer.Route(s)

    // State
    s = r.PathPrefix("/state").Subrouter()
    state.Route(s)

    // Band
    s = r.PathPrefix("/band").Subrouter()
    band.Route(s)

    // Ticket
    s = r.PathPrefix("/ticket").Subrouter()
    ticket.Route(s)

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    http.ListenAndServe(":8080", nil)
}

func HandleHome(rw http.ResponseWriter, req *http.Request) {
    home.GetPage(rw, req)
}