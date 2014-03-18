package state

import (
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/common"
)

func Route(s *mux.Router) {
    s.HandleFunc("/", viewAllHandler)
    s.HandleFunc("/{id:[0-9]+}{_:/?}", viewOneHandler)
    s.HandleFunc("/view/{id:[0-9]+}", viewOneHandler)
    s.HandleFunc("/edit/{id:[0-9]+}", editHandler)
    s.HandleFunc("/add{_:/?}", addHandler)
}

func viewAllHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "state",
        Title:      "View All Controller",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["state.html"] = template.Must(template.ParseFiles("templates/state/viewAll.html", "templates/index.html"))
    tmpl["state.html"].ExecuteTemplate(rw, "base", p)
}

func viewOneHandler(rw http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id := params["id"]

    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "state",
        Title:      "View One Controller: "+id,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["state.html"] = template.Must(template.ParseFiles("templates/state/viewOne.html", "templates/index.html"))
    tmpl["state.html"].ExecuteTemplate(rw, "base", p)
}

func editHandler(rw http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id := params["id"]

    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "state",
        Title:      "Edit Controller: "+id,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["state.html"] = template.Must(template.ParseFiles("templates/state/edit.html", "templates/index.html"))
    tmpl["state.html"].ExecuteTemplate(rw, "base", p)
}

func addHandler(rw http.ResponseWriter, req *http.Request) {
    if req.Method == "POST" {
        dbmap := db.InitDb()
        defer dbmap.Db.Close()

        err := req.ParseForm()
        common.CheckError(err)
        form := req.Form

        err = dbmap.Insert(&db.State{Name: form["name"][0], Acronym: form["acronym"][0]})
        common.CheckError(err)
    }

    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "state",
        Title:      "Add Controller",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["state.html"] = template.Must(template.ParseFiles("templates/state/add.html", "templates/index.html"))
    tmpl["state.html"].ExecuteTemplate(rw, "base", p)
}