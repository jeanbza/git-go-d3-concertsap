package retailer

import (
    "net/http"
    "html/template"
    "github.com/gorilla/mux"
    "git-go-d3-concertsap/app/state"
    "git-go-d3-concertsap/app/database"
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
        PageName:   "retailer",
        Title:      "View All Controller",
    }

    tmpl := make(map[string]*template.Template)
    tmpl["retailer.html"] = template.Must(template.ParseFiles("templates/retailer/viewAll.html", "templates/index.html"))
    tmpl["retailer.html"].ExecuteTemplate(rw, "base", p)
}

func viewOneHandler(rw http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id := params["id"]

    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "retailer",
        Title:      "View One Controller: "+id,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["retailer.html"] = template.Must(template.ParseFiles("templates/retailer/viewOne.html", "templates/index.html"))
    tmpl["retailer.html"].ExecuteTemplate(rw, "base", p)
}

func editHandler(rw http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id := params["id"]

    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "retailer",
        Title:      "Edit Controller: "+id,
    }

    tmpl := make(map[string]*template.Template)
    tmpl["retailer.html"] = template.Must(template.ParseFiles("templates/retailer/edit.html", "templates/index.html"))
    tmpl["retailer.html"].ExecuteTemplate(rw, "base", p)
}

func addHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
        States      []db.State
    }

    p := Page{
        PageName:   "retailer",
        Title:      "Add Controller",
        States:     state.FindAll(),
    }

    tmpl := make(map[string]*template.Template)
    tmpl["retailer.html"] = template.Must(template.ParseFiles("templates/retailer/add.html", "templates/index.html"))
    tmpl["retailer.html"].ExecuteTemplate(rw, "base", p)
}