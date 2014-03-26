package retailer

import (
    "net/http"
    "html/template"

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

    retailer := Retailer{
        Name:       common.IssetInForm(form["name"], 0),
        Website:    common.IssetInForm(form["website"], 0),
    }

    insertRetailer(retailer)

    http.Redirect(rw, req, "add", http.StatusFound)
}

func viewAllHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
        Retailers   []Retailer
    }

    retailers := FindAll()

    p := Page{
        PageName:   "retailer",
        Title:      "View All Retailers",
        Retailers:  retailers,
    }

    common.Templates = template.Must(template.ParseFiles("templates/retailer/viewAll.html", common.LayoutPath))
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
        PageName:   "retailer",
        Title:      "View One Controller: "+id,
    }

    common.Templates = template.Must(template.ParseFiles("templates/retailer/viewOne.html", common.LayoutPath))
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
        PageName:   "retailer",
        Title:      "Edit Controller: "+id,
    }

    common.Templates = template.Must(template.ParseFiles("templates/retailer/edit.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func addHandler(rw http.ResponseWriter, req *http.Request) {
    if req.Method == "POST" {
        dbmap := db.InitDb(Retailer{}, "retailer")
        defer dbmap.Db.Close()

        err := req.ParseForm()
        common.CheckError(err)
        form := req.Form

        err = dbmap.Insert(&Retailer{Name: form["name"][0], Website: form["website"][0]})
        common.CheckError(err)
    }

    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "retailer",
        Title:      "Add Controller",
    }

    common.Templates = template.Must(template.ParseFiles("templates/retailer/add.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}