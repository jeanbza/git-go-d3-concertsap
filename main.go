package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "git-go-d3-concertsap/src/data"
)

func main() {
    r := mux.NewRouter()

    http.Handle("/", r)
    r.HandleFunc("/input/{form:[a-zA-Z]*}", HandleInput).Methods("GET")

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