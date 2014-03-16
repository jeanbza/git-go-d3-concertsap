package collect

import (
    "net/http"
    "html/template"
    "git-go-d3-concertsap/src/database"
    "git-go-d3-concertsap/src/common"
)

func getCollectFiles(args ...string) ([]string) {
    files := []string{"templates/collect/collect.html", "templates/index.html"}
    
    for i := range args {
        files = append(files, args[i])
    }
    
    return files
}

func GetConcertForm(rw http.ResponseWriter, req *http.Request) {
    type State struct {
        Id   int
        Name string
    }

    type Page struct {
        PageName    string
        Title       string
        States      []State
    }

    states := []State{}

    var (
        id      int
        name    string
    )

    rows := database.Select("SELECT id, name FROM state ORDER BY name", "concertsap")

    for rows.Next() {
        err := rows.Scan(&id, &name)
        common.CheckError(err)

        states = append(states, State{Id: id, Name: name})
    }

    p := Page{
        PageName:   "data",
        Title:      "Add A Concert",
        States:     states,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["data.html"] = template.Must(template.ParseFiles(getCollectFiles("templates/collect/concertform.html")...))
    tmpl["data.html"].ExecuteTemplate(rw, "base", p)
}

func GetTicketForm(rw http.ResponseWriter, req *http.Request) {
    /*
     * define concert and retailers arrays to pass into
     * Page instantiation
     */
    type Concert struct {
        Id      int
        Name    string
    }

    type Retailer struct {
        Id      int
        Name    string
    }

    type Page struct {
        PageName    string
        Title       string
        Concerts    []Concert
        Retailers   []Retailer 
    }

    concerts := []Concert{}
    retailers := []Retailer{}

    var (
        id      int
        name    string
    )

    // connect to db and query concerts
    rows := database.Select("SELECT id,name FROM concert ORDER BY name", "concertsap")

    for rows.Next() {
        err := rows.Scan(&id, &name)
        common.CheckError(err)
        concerts = append(concerts, Concert{ Id:id, Name:name })
    }

    // connect to db and query retailers
    rows = database.Select("SELECT id,name FROM retailer ORDER BY name", "concertsap")

    for rows.Next() {
        err := rows.Scan(&id, &name)
        common.CheckError(err)
        retailers = append(retailers, Retailer{ Id:id, Name:name })
    }

    p := Page{
        PageName:   "data",
        Title:      "Add A Ticket Record",
        Concerts:   concerts,
        Retailers:  retailers,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["data.html"] = template.Must(template.ParseFiles(getCollectFiles("templates/collect/ticketform.html")...))
    tmpl["data.html"].ExecuteTemplate(rw, "base", p)
}

func Get404(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        Title string
    }

    p := Page{
        Title: "data",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["data.html"] = template.Must(template.ParseFiles(getCollectFiles("templates/collect/404.html")...))
    tmpl["data.html"].ExecuteTemplate(rw, "base", p)
}