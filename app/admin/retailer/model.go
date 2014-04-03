package retailer

import (
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/common"
)

type Retailer struct {
    Id      int64   `db:"id"`
    Name    string  `db:"name"`
    Website string  `db:"website"`
}

func insertRetailer(retailer Retailer) {
    dbmap := db.InitDb(Retailer{}, "retailer")
    defer dbmap.Db.Close()

    dbmap.Insert(&retailer)
}

func FindAll() ([]Retailer) {
    dbmap := db.InitDb(Retailer{}, "retailer")
    defer dbmap.Db.Close()

    var retailers []Retailer
    _, err := dbmap.Select(&retailers, "SELECT * FROM retailer ORDER BY name")
    common.CheckError(err)

    return retailers
}