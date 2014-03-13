package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "html/template"
    "net/http"
    "os"
    "fmt"
    "log"
)

type Page struct {
    Title string
    Posts []Post
}

type Post struct {
    Title   string
    Content string
}

func main() {
    http.HandleFunc("/", homePage)

    fileServer := http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
    http.Handle("/css/", fileServer)

    fileServer = http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
    http.Handle("/js/", fileServer)

    fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
    http.Handle("/html/", fileServer)

    fmt.Println("begin")

    // err := http.ListenAndServe(":8080", nil)
    // checkError(err)

    fmt.Println("1")

    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/whiteboard")
    checkError(err)

    fmt.Println("2")
    
    var (
        id int
        note string
    )

    rows, err := db.Query("SELECT id, note FROM load_records")
    checkError(err)
    defer rows.Close()

    for rows.Next() {
        err := rows.Scan(&id, &note)
        log.Println(id, note)
        checkError(err)
    }

    err = rows.Err()
    checkError(err)

    fmt.Println("4")

}

func homePage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "home",
        Posts: nil,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["contact.html"] = template.Must(template.ParseFiles("html/home.html", "html/index.html"))
    tmpl["contact.html"].ExecuteTemplate(rw, "base", p)
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}