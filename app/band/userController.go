package band

import (
    "net/http"
    "html/template"

    "git-go-d3-concertsap/app/common"
    "git-go-d3-concertsap/app/concert"

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
        Bands       []BandAndConcertsCount
    }

    bandsWithConcertsCount := FindBandsAndConcerts()

    p := Page{
        PageName:   "user_band",
        Title:      "View All Bands",
        Bands:      bandsWithConcertsCount,
    }

    common.Templates = template.Must(template.ParseFiles("templates/band/userViewAll.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func userViewOneHandler(rw http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id := params["id"]

    type Page struct {
        PageName    string
        Title       string
        Band        Band
        Concerts    []concert.Concert
    }

    band := FindOne(id)
    concerts := concert.FindConcertsByBand(id)

    p := Page{
        PageName:   "user_band",
        Title:      "View One Controller: "+id,
        Band:       band,
        Concerts:   concerts,
    }

    common.Templates = template.Must(template.ParseFiles("templates/band/userViewOne.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}