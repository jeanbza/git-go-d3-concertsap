package band

import (
    "git-go-d3-concertsap/app/database"
)

type BandConcert struct {
    Id          int64   `db:"id"`
    BandId      int64   `db:"band_id"`
    ConcertId   int64   `db:"concert_id"`
}

func insertBandConcert(bandConcert BandConcert) {
    dbmap := db.InitDb(BandConcert{}, "band_concert")
    defer dbmap.Db.Close()

    dbmap.Insert(&bandConcert)
}