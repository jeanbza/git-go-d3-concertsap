package concert

import (
    "net/http"
    "html/template"
    "strconv"
    "time"
    
    "git-go-d3-concertsap/app/common"
    "git-go-d3-concertsap/app/state"

    "github.com/gorilla/mux"
)

func RouteAdmin(s *mux.Router) {
    s.HandleFunc("/", adminViewAllHandler)
    s.HandleFunc("/{id:[0-9]+}{_:/?}", adminViewOneHandler)
    s.HandleFunc("/view/{id:[0-9]+}", adminViewOneHandler)
    s.HandleFunc("/edit/{id:[0-9]+}", editHandler)
    s.HandleFunc("/add{_:/?}", addHandler)
    s.HandleFunc("/save{_:/?}", saveHandler).Methods("POST")
}

func adminViewAllHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
        Concerts    []Concert
    }

    concerts := FindAll()

    p := Page{
        PageName:   "admin_concert",
        Title:      "View All Concerts",
        Concerts:   concerts,
    }

    common.Templates = template.Must(template.ParseFiles("templates/concert/adminViewAll.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func adminViewOneHandler(rw http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id := params["id"]

    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "admin_concert",
        Title:      "View One Controller: "+id,
    }

    common.Templates = template.Must(template.ParseFiles("templates/concert/adminViewOne.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func saveHandler(rw http.ResponseWriter, req *http.Request) {
    err := req.ParseForm()
    common.CheckError(err)
    form := req.Form

    stateId, err := strconv.ParseInt(common.IssetInForm(form["state_id"], 0), 10, 64)
    common.CheckError(err)

    concert := Concert{
        Name:       common.IssetInForm(form["name"], 0),
        Address:    common.IssetInForm(form["address"], 0),
        StateId:    stateId,
        Website:    common.IssetInForm(form["website"], 0),
        Start:      common.IssetInForm(form["start"], 0),
        End:        common.IssetInForm(form["end"], 0),
    }

    insertConcert(concert)

    http.Redirect(rw, req, "add", http.StatusFound)
}

func editHandler(rw http.ResponseWriter, req *http.Request) {
    params := mux.Vars(req)
    id := params["id"]

    type Page struct {
        PageName    string
        Title       string
    }

    p := Page{
        PageName:   "admin_concert",
        Title:      "Edit Controller: "+id,
    }

    common.Templates = template.Must(template.ParseFiles("templates/concert/edit.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func addHandler(rw http.ResponseWriter, req *http.Request) {
    type Page struct {
        PageName    string
        Title       string
        States      []state.State
        Now         string
    }

    t := time.Now()

    p := Page{
        PageName:   "admin_concert",
        Title:      "Add Controller",
        States:     state.FindAll(),
        Now:        t.Format("2006-01-02"),
    }

    common.Templates = template.Must(template.ParseFiles("templates/concert/add.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}