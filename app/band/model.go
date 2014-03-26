package band

import (
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/common"
)

type Band struct {
    Id      int64   `db:"id"`
    Name    string  `db:"name"`
    Website string  `db:"website"`
}

func insertBand(band Band) {
    dbmap := db.InitDb(Band{}, "band")
    defer dbmap.Db.Close()

    dbmap.Insert(&band)
}

func FindAll() ([]Band) {
    dbmap := db.InitDb(Band{}, "band")
    defer dbmap.Db.Close()

    var bands []Band
    _, err := dbmap.Select(&bands, "SELECT * FROM band ORDER BY name")
    common.CheckError(err)

    return bands
}