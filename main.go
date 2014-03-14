package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "git-go-d3-concertsap/src/data"
    "git-go-d3-concertsap/src/home"
    "git-go-d3-concertsap/src/sqlEntry"
)

func main() {
    r := mux.NewRouter()

    http.Handle("/", r)
    
    r.HandleFunc("/", HandleHome).Methods("GET")
    r.HandleFunc("/input/{form:[a-zA-Z]*}", HandleInput).Methods("GET")
    r.HandleFunc("/sqlEntry", sqlEntryHandler).Methods("GET")

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    http.ListenAndServe(":8080", nil)
}

func HandleInput(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    form := params["form"]

    switch form {
    case "concert":
        data.GetConcertForm(w, r)
    case "ticket":
        data.GetTicketForm(w, r)
    default:
        data.Get404(w, r)
    }
}

func HandleHome(rw http.ResponseWriter, req *http.Request) {
    home.GetPage(rw, req)
}

func sqlEntryHandler(rw http.ResponseWriter, req *http.Request) {
    sqlEntry.GetPage(rw, req)
}