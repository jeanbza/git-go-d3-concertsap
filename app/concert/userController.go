package concert

import (
    "net/http"
    "html/template"
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
        PageName        string
        Title           string
        Concert         Concert
        ConcertState    int
        UntilDays       int
    }

    concert := FindOne(id)

    var concertState int
    var untilDays int

    tStart, terr := time.Parse("2006-01-02 15:04:05", concert.Start+" 23:59:59")
    common.CheckError(terr)

    tEnd, terr := time.Parse("2006-01-02 15:04:05", concert.End+" 23:59:59")
    common.CheckError(terr)

    if (time.Now().After(tEnd)) {
        // Concert is in the past - concertState = 1
        concertState = 1
        untilDays = int(time.Since(tEnd).Hours()/24)+1
    } else if (time.Now().After(tStart)) {
        // Concert is now! - concertState = 2
        concertState = 2
        untilDays = int(time.Since(tEnd).Hours()/24)+1
    } else {
        // Concert is in the future - concertState = 3
        concertState = 3
        untilDays = -1*int(time.Since(tStart).Hours()/24)+1
    }

    p := Page{
        PageName:       "user_concert",
        Title:          "View One Controller: "+id,
        Concert:        concert,
        ConcertState:   concertState,
        UntilDays:      untilDays,
    }

    common.Templates = template.Must(template.ParseFiles("templates/concert/userViewOne.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}