package main

import (
    "github.com/gorilla/mux"
    "log"
    "net/http"
)

func main() {
    r := mux.NewRouter()

    http.Handle("/", r)
    r.HandleFunc("/input/{form:[a-zA-Z]+}", inputForm).Methods("GET")

    log.Println("Listening...")
    http.ListenAndServe(":8080", nil)
}

func inputForm(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    name := params["form"]
    w.Write([]byte("Form " + name))
}