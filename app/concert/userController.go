package concert

import (
    "net/http"
    "html/template"
    
    "git-go-d3-concertsap/app/common"

    "github.com/gorilla/mux"
)

func RouteUser(s *mux.Router) {
    s.HandleFunc("/", userViewAllHandler)
    s.HandleFunc("/{id:[0-9]+}{_:/?}", userViewOneHandler)
    s.HandleFunc("/view/{id:[0-9]+}", userViewOneHandler)
}

func userViewAllHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
        Concerts    []Concert
    }

    concerts := FindAll()

    p := Page{
        PageName:   "user_concert",
        Title:      "View All Concerts",
        Concerts:   concerts,
    }

    common.Templates = template.Must(template.ParseFiles("templates/concert/userViewAll.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func userViewOneHandler(rw http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id := params["id"]

    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "user_concert",
        Title:      "View One Controller: "+id,
    }

    common.Templates = template.Must(template.ParseFiles("templates/concert/userViewOne.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}