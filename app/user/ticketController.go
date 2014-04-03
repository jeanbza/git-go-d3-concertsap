package ticket

import (
    "net/http"
    "html/template"
    "time"

    "git-go-d3-concertsap/app/common"

    "github.com/gorilla/mux"
)

func Route(s *mux.Router) {
    s.HandleFunc("/", addHandler)
    s.HandleFunc("/add{_:/?}", addHandler)
    s.HandleFunc("/save{_:/?}", saveHandler).Methods("POST")
}

func viewAllHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
    }

    concerts := FindAll()

    p := Page{
        PageName:   "concert",
        Title:      "View All Concerts",
        Concerts:   concerts,
    }

    common.Templates = template.Must(template.ParseFiles("templates/admin/concert/viewAll.html", common.LayoutPath))
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
        PageName:   "concert",
        Title:      "View One Controller: "+id,
    }

    common.Templates = template.Must(template.ParseFiles("templates/admin/concert/viewOne.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}