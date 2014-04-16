package band

import (
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/common"
)

type BandConcert struct {
    Id          int64   `db:"id"`
    BandId      int64   `db:"band_id"`
    ConcertId   int64   `db:"concert_id"`
}

type BandJoinedConcert struct {
    BandId          int64   `db:"band_id"`
    BandName        string  `db:"band_name"`
    BandWebsite     string  `db:"band_website"`
    ConcertId       int64   `db:"concert_id"`
    ConcertName     string  `db:"concert_name"`
    ConcertAddress  string  `db:"concert_address"`
    ConcertStateId  int64   `db:"concert_state_id"`
    ConcertWebsite  string  `db:"concert_website"`
    ConcertStart    string  `db:"concert_start"`
    ConcertEnd      string  `db:"concert_end"`
}

func insertBandConcert(bandConcert BandConcert) {
    dbmap := db.InitDb(BandConcert{}, "band_concert")
    defer dbmap.Db.Close()

    dbmap.Insert(&bandConcert)
}

func FindBandConcert(concertId string) ([]BandJoinedConcert) {
    dbmap := db.InitDb(BandConcert{}, "band_concert")
    defer dbmap.Db.Close()

    var concertIdString string

    if (concertId != "") {
        concertIdString = "AND c.id = "+concertId
    }

    var bandConcerts []BandJoinedConcert
    _, err := dbmap.Select(&bandConcerts, `
        SELECT b.id AS band_id,
            b.name AS band_name,
            b.website AS band_website,
            c.id AS concert_id,
            c.name AS concert_name,
            c.address AS concert_address,
            c.state_id AS concert_state_id,
            c.website AS concert_website,
            c.start AS concert_start,
            c.end AS concert_end
        FROM band_concert bc
        JOIN band b
        ON b.id = bc.band_id
        JOIN concert c
        ON c.id = bc.concert_id
        `+concertIdString+`
        ORDER BY b.name`)
    common.CheckError(err)

    return bandConcerts
}