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

type BandAndConcertsCount struct {
    Id              int64   `db:"id"`
    Name            string  `db:"name"`
    ConcertsCount   int `db:"count_count"`   
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

func FindBandsAndConcerts() ([]BandAndConcertsCount) {
    dbmap := db.InitDb(Band{}, "band")
    defer dbmap.Db.Close()

    var bandsAndConcertsCount []BandAndConcertsCount
    _, err := dbmap.Select(&bandsAndConcertsCount, "SELECT b.id, b.name, COUNT(*) AS count_count FROM band_concert bc JOIN band b ON bc.band_id = b.id GROUP BY b.id ORDER BY b.name")
    common.CheckError(err)

    return bandsAndConcertsCount   
}