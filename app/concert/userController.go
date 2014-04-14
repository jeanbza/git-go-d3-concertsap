package concert

import (
    "net/http"
    "html/template"
    "log"
    "time"
    
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
        Concert     Concert
    }

    concert := FindOne(id)

    // var concertInFuture bool
    // var untilDays int

    tStart, terr := time.Parse("2006-01-02", "2014-04-08")
    common.CheckError(terr)

    tEnd, terr := time.Parse("2006-01-02", "2014-04-26")
    common.CheckError(terr)

    if (time.Now().After(tStart)) {
        // Concert is in the future
        // concertInFuture = true

        log.Println("start")
        log.Println(int(time.Since(tStart).Hours()/24))
    } else {
        // Concert is in the past
        
        log.Println("end")
        log.Println(time.Since(tEnd).Hours()/24)
    }

    p := Page{
        PageName:   "user_concert",
        Title:      "View One Controller: "+id,
        Concert:    concert,
    }

    common.Templates = template.Must(template.ParseFiles("templates/concert/userViewOne.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}