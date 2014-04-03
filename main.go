package main

import (
    "net/http"

    "git-go-d3-concertsap/app/home"
    "git-go-d3-concertsap/app/admin/concert"
    "git-go-d3-concertsap/app/admin/state"
    "git-go-d3-concertsap/app/admin/band"
    "git-go-d3-concertsap/app/admin/ticket"
    "git-go-d3-concertsap/app/admin/retailer"

    "github.com/gorilla/mux"
)

func main() {
    r := mux.NewRouter()

    http.Handle("/", r)
    
    r.HandleFunc("/", HandleHome).Methods("GET")

    // Concert
    s := r.PathPrefix("/admin/concert").Subrouter()
    concert.Route(s)

    // Retailer
    s = r.PathPrefix("/admin/retailer").Subrouter()
    retailer.Route(s)

    // State
    s = r.PathPrefix("/admin/state").Subrouter()
    state.Route(s)

    // Band
    s = r.PathPrefix("/admin/band").Subrouter()
    band.Route(s)

    // Ticket
    s = r.PathPrefix("/admin/ticket").Subrouter()
    ticket.Route(s)

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    http.ListenAndServe(":8080", nil)
}

func HandleHome(rw http.ResponseWriter, req *http.Request) {
    home.GetPage(rw, req)
}