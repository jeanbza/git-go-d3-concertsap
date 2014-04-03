package band

import (
    "net/http"
    "html/template"
    "regexp"
    "strconv"
    "strings"

    "git-go-d3-concertsap/app/common"
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/admin/concert"

    "github.com/gorilla/mux"
)

func Route(s *mux.Router) {
    s.HandleFunc("/", viewAllHandler)
    s.HandleFunc("/{id:[0-9]+}{_:/?}", viewOneHandler)
    s.HandleFunc("/view/{id:[0-9]+}", viewOneHandler)
    s.HandleFunc("/edit/{id:[0-9]+}", editHandler)
    s.HandleFunc("/add{_:/?}", addHandler)
    s.HandleFunc("/addBandsToConcert{_:/?}", addBandsToConcertHandler)
    s.HandleFunc("/addBandRecordToConcert{_:/?}", addBandRecordToConcertHandler)
    s.HandleFunc("/save{_:/?}", saveHandler).Methods("POST")
    s.HandleFunc("/saveBandsToConcert{_:/?}", saveBandsToConcertHandler).Methods("POST")
    s.HandleFunc("/saveBandRecordToConcert{_:/?}", saveBandRecordToConcertHandler).Methods("POST")
}

func saveHandler(rw http.ResponseWriter, req *http.Request) {
    dbmap := db.InitDb(Band{}, "band")
    defer dbmap.Db.Close()

    err := req.ParseForm()
    common.CheckError(err)
    form := req.Form

    newBand := Band{Name: strings.ToUpper(form["name"][0]), Website: form["website"][0]}
    insertBand(newBand)

    http.Redirect(rw, req, "add", http.StatusFound)
}

func saveBandsToConcertHandler(rw http.ResponseWriter, req *http.Request) {
    dbmap := db.InitDb(Band{}, "band")
    defer dbmap.Db.Close()

    err := req.ParseForm()
    common.CheckError(err)
    form := req.Form

    concertId, err := strconv.ParseInt(form["concert_id"][0], 10, 64)
    common.CheckError(err)

    r, err := regexp.Compile(`([^<]+)<br>`)
    common.CheckError(err)
    rawBands := form["parsedInput"][0]
    bandMatches := r.FindAllStringSubmatch(rawBands, -1)
    
    for _,bandName := range bandMatches {
        newBand := Band{Name: bandName[1]}
        err = dbmap.SelectOne(&newBand, "SELECT * FROM concertsap.band WHERE name=?", strings.ToUpper(newBand.Name))

        if err != nil {
            insertBand(newBand)
            err = dbmap.SelectOne(&newBand, "SELECT * FROM concertsap.band WHERE name=?", newBand.Name)
        }

        newBandConcert := BandConcert{BandId: newBand.Id, ConcertId: concertId}
        insertBandConcert(newBandConcert)
    }

    http.Redirect(rw, req, "addBandsToConcert", http.StatusFound)   
}

func saveBandRecordToConcertHandler(rw http.ResponseWriter, req *http.Request) {
    dbmap := db.InitDb(Band{}, "band")

    err := req.ParseForm()
    common.CheckError(err)
    form := req.Form

    concertId, err := strconv.ParseInt(form["concert_id"][0], 10, 64)
    common.CheckError(err)
    newBand := Band{Name: strings.ToUpper(form["name"][0])}
    err = dbmap.SelectOne(&newBand, "SELECT * FROM concertsap.band WHERE name=?", strings.ToUpper(newBand.Name))

    if err != nil {
        insertBand(newBand)
        err = dbmap.SelectOne(&newBand, "SELECT * FROM concertsap.band WHERE name=?", newBand.Name)
    }

    dbmap.Db.Close()
    dbmap = db.InitDb(BandConcert{}, "band_concert")
    defer dbmap.Db.Close()

    bandConcert := BandConcert{BandId: newBand.Id, ConcertId: concertId}
    err = dbmap.SelectOne(&bandConcert, "SELECT * FROM concertsap.concert_band WHERE band_id=? AND concert_id=?", bandConcert.BandId, bandConcert.ConcertId)

    if err != nil {
        insertBandConcert(bandConcert)
        err = dbmap.SelectOne(&bandConcert, "SELECT * FROM concertsap.concert_band WHERE band_id=? AND concert_id=?", bandConcert.BandId, bandConcert.ConcertId)
    }

    formDate := form["date"][0]
    formTime := form["time"][0]
    bandConcertRecord := BandConcertRecord{BandId: bandConcert.BandId, ConcertId: bandConcert.ConcertId, Date: formDate, Time: formTime}
    insertBandConcertRecord(bandConcertRecord)

    http.Redirect(rw, req, "addBandRecordToConcert", http.StatusFound)   
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

    common.Templates = template.Must(template.ParseFiles("templates/admin/band/viewAll.html", common.LayoutPath))
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

    common.Templates = template.Must(template.ParseFiles("templates/admin/band/viewOne.html", common.LayoutPath))
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

    common.Templates = template.Must(template.ParseFiles("templates/admin/band/edit.html", common.LayoutPath))
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

    common.Templates = template.Must(template.ParseFiles("templates/admin/band/add.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func addBandsToConcertHandler(rw http.ResponseWriter, req *http.Request) {
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

    common.Templates = template.Must(template.ParseFiles("templates/admin/band/addBandsToConcert.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}

func addBandRecordToConcertHandler(rw http.ResponseWriter, req *http.Request) {
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

    common.Templates = template.Must(template.ParseFiles("templates/admin/band/addBandRecordToConcert.html", common.LayoutPath))
    err := common.Templates.ExecuteTemplate(rw, "base", p)
    common.CheckError(err)
}