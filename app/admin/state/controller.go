package state

import (
    "net/http"
    "html/template"
    //"strconv"
    
    "git-go-d3-concertsap/app/common"
    "git-go-d3-concertsap/app/database"

    "github.com/gorilla/mux"
)

func Route(s *mux.Router) {
    s.HandleFunc("/", viewAllHandler)
    s.HandleFunc("/{id:[0-9]+}{_:/?}", viewOneHandler)
    s.HandleFunc("/view/{id:[0-9]+}", viewOneHandler)
    s.HandleFunc("/edit/{id:[0-9]+}", editHandler)
    s.HandleFunc("/add{_:/?}", addHandler)
    s.HandleFunc("/save{_:/?}", saveHandler).Methods("POST")
}

func saveHandler(rw http.ResponseWriter, req *http.Request) {
    err := req.ParseForm()
    common.CheckError(err)
    form := req.Form

    state := State{
        Name:       common.IssetInForm(form["name"], 0),
        Acronym:    common.IssetInForm(form["acronym"], 0),
    }

    insertState(state)

    http.Redirect(rw, req, "add", http.StatusFound)
}

func viewAllHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
        States      []State
    }

    states := FindAll()

    p := Page{
        PageName:   "state",
        Title:      "View All States",
        States:     states,
    }

    common.Templates = template.Must(template.ParseFiles("templates/admin/state/viewAll.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
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

    common.Templates = template.Must(template.ParseFiles("templates/admin/state/viewOne.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
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

    common.Templates = template.Must(template.ParseFiles("templates/admin/state/edit.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func addHandler(rw http.ResponseWriter, req *http.Request) {
    if req.Method == "POST" {
        dbmap := db.InitDb(State{}, "state")
        defer dbmap.Db.Close()

        err := req.ParseForm()
        common.CheckError(err)
        form := req.Form

        err = dbmap.Insert(&State{Name: form["name"][0], Acronym: form["acronym"][0]})
        common.CheckError(err)
    }

    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "state",
        Title:      "Add New Concert",
    }

    common.Templates = template.Must(template.ParseFiles("templates/admin/state/add.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}