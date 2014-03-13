package main

import (
    "html/template"
    "net/http"
    "os"
    "fmt"
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
    http.HandleFunc("/", blogPage)
    http.HandleFunc("/blog", blogPage)
    http.HandleFunc("/blog/", blogPage)

    http.HandleFunc("/about", aboutPage)
    http.HandleFunc("/about/", aboutPage)

    http.HandleFunc("/contact", contactPage)
    http.HandleFunc("/contact/", contactPage)

    fileServer := http.StripPrefix("/css/", http.FileServer(http.Dir("css")))
    http.Handle("/css/", fileServer)

    fileServer = http.StripPrefix("/js/", http.FileServer(http.Dir("js")))
    http.Handle("/js/", fileServer)

    fileServer = http.StripPrefix("/html/", http.FileServer(http.Dir("html")))
    http.Handle("/html/", fileServer)

    err := http.ListenAndServe(":8080", nil)
    checkError(err)
}

func blogPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "blog",
        Posts: []Post{
            Post{
                Title: "Test blog title 1",
                Content: "Test blog content 1",
            },
            Post{
                Title: "Test blog title 2",
                Content: "Test blog content 2",
            },
            Post{
                Title: "Test blog title THE LAST",
                Content: "Test blog content BOOM THE LAST",
            },
        },
    }

    tmpl := make(map[string]*template.Template)
    tmpl["blog.html"] = template.Must(template.ParseFiles("html/blog.html", "html/index.html"))
    tmpl["blog.html"].ExecuteTemplate(rw, "base", p)
}

func aboutPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "about",
        Posts: nil,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["about.html"] = template.Must(template.ParseFiles("html/about.html", "html/index.html"))
    tmpl["about.html"].ExecuteTemplate(rw, "base", p)
}

func contactPage(rw http.ResponseWriter, req *http.Request) {
    p := Page{
        Title: "contact",
        Posts: nil,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["contact.html"] = template.Must(template.ParseFiles("html/contact.html", "html/index.html"))
    tmpl["contact.html"].ExecuteTemplate(rw, "base", p)
}

func checkError(err error) {
    if err != nil {
        fmt.Println("Fatal error ", err.Error())
        os.Exit(1)
    }
}