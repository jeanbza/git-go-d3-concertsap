package band

import (
    "net/http"
    "html/template"

    "git-go-d3-concertsap/app/common"
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/concert"

    "github.com/gorilla/mux"
)

func Route(s *mux.Router) {
    s.HandleFunc("/", viewAllHandler)
    s.HandleFunc("/{id:[0-9]+}{_:/?}", viewOneHandler)
    s.HandleFunc("/view/{id:[0-9]+}", viewOneHandler)
    s.HandleFunc("/edit/{id:[0-9]+}", editHandler)
    s.HandleFunc("/add{_:/?}", addHandler)
    s.HandleFunc("/addConcertBand{_:/?}", addConcertBandHandler)
    s.HandleFunc("/save{_:/?}", saveHandler).Methods("POST")
}

func saveHandler(rw http.ResponseWriter, req *http.Request) {
    dbmap := db.InitDb(Band{}, "band")
    defer dbmap.Db.Close()

    err := req.ParseForm()
    common.CheckError(err)
    form := req.Form

    err = dbmap.Insert(&Band{Name: form["name"][0], Website: form["website"][0]})
    common.CheckError(err)

    http.Redirect(rw, req, "add", http.StatusFound)
}

func viewAllHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
        Bands       []Band
    }

    bands := FindAll()

    p := Page{
        PageName:   "band",
        Title:      "View All Bands",
        Bands:      bands,
    }

    common.Templates = template.Must(template.ParseFiles("templates/band/viewAll.html", common.LayoutPath))
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
        PageName:   "band",
        Title:      "View One Controller: "+id,
    }

    common.Templates = template.Must(template.ParseFiles("templates/band/viewOne.html", common.LayoutPath))
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
        PageName:   "band",
        Title:      "Edit Controller: "+id,
    }

    common.Templates = template.Must(template.ParseFiles("templates/band/edit.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func addHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "band",
        Title:      "Add Controller",
    }

    common.Templates = template.Must(template.ParseFiles("templates/band/add.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func addConcertBandHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
        Concerts    []concert.Concert
    }

    concerts := concert.FindAll()

    p := Page{
        PageName:   "band",
        Title:      "Add Controller",
        Concerts:   concerts,
    }

    common.Templates = template.Must(template.ParseFiles("templates/band/addConcertBand.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}