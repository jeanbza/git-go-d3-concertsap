package band

import (
    "git-go-d3-concertsap/app/database"
)

type BandConcertRecord struct {
    Id          int64   `db:"id"`
    BandId      int64   `db:"band_id"`
    ConcertId   int64   `db:"concert_id"`
    Date        int64   `db:"date"`
    Time        int64   `db:"time"`
}

func insertBandConcertRecord(bandConcertRecord BandConcertRecord) {
    dbmap := db.InitDb(BandConcertRecord{}, "band_concert_record")
    defer dbmap.Db.Close()

    dbmap.Insert(&bandConcertRecord)
}