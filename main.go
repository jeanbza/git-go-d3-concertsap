package main

import (
    "github.com/gorilla/mux"
    "net/http"
    "git-go-d3-concertsap/src/collect"
    "git-go-d3-concertsap/src/home"
)

func main() {
    r := mux.NewRouter()

    http.Handle("/", r)
    
    r.HandleFunc("/", HandleHome).Methods("GET")
    r.HandleFunc("/collect/{form:[a-zA-Z]*}", HandleCollect).Methods("GET")

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    http.ListenAndServe(":8080", nil)
}

func HandleCollect(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    form := params["form"]

    switch form {
    case "concert":
        collect.GetConcertForm(w, r)
    case "ticket":
        collect.GetTicketForm(w, r)
    case "save":
        collect.SaveForm(w, r)
    default:
        collect.Get404(w, r)
    }
}

func HandleHome(rw http.ResponseWriter, req *http.Request) {
    home.GetPage(rw, req)
}