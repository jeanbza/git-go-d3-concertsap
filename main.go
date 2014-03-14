package main

import (
    "git-go-d3-concertsap/src/common"
//    "git-go-d3-concertsap/src/database"
    "git-go-d3-concertsap/src/home"
    "git-go-d3-concertsap/src/sqlEntry"
    "net/http"
)

func main() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/sqlEntry", sqlPage)
    http.HandleFunc("/sqlEntrySave", sqlPageSave)

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    //database.Insert("INSERT INTO state (name, acronym) VALUES ('Virginia', 'VA'), ('Washington', 'WA'), ('Arizona', 'AZ'), ('Colorado', 'CO')", "concertsap")

    err := http.ListenAndServe(":8080", nil)
    common.CheckError(err)
}

func homePage(rw http.ResponseWriter, req *http.Request) {
    home.GetPage(rw, req)
}

func sqlPage(rw http.ResponseWriter, req *http.Request) {
    sqlEntry.GetPage(rw, req)
}

func sqlPageSave(rw http.ResponseWriter, req *http.Request) {
    sqlEntry.Save(rw, req)
}