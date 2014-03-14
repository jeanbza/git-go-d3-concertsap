package main

import (
    "git-go-d3-concertsap/src/home"
    "git-go-d3-concertsap/src/common"
    "git-go-d3-concertsap/src/database"
    "log"
    "net/http"
)

func main() {
    http.HandleFunc("/", homePage)

    fileServer := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
    http.Handle("/static/", fileServer)

    var (
        id int
        acronym string
    )

    rows := database.Select("SELECT id,acronym FROM state", "concertsap")

    for rows.Next() {
        err := rows.Scan(&id, &acronym)
        log.Println(id, acronym)
        common.CheckError(err)
    }

    // err := http.ListenAndServe(":8080", nil)
    // common.CheckError(err)
}

func homePage(rw http.ResponseWriter, req *http.Request) {
    home.GetPage(rw, req)
}
