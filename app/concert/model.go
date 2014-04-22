package concert

import (
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/common"
)

type Concert struct {
    Id      int64   `db:"id"`
    Name    string  `db:"name"`
    Address string  `db:"address"`
    StateId int64   `db:"state_id"`
    Website string  `db:"website"`
    Start   string  `db:"start"`
    End     string  `db:"end"`
}

func insertConcert(concert Concert) {
    dbmap := db.InitDb(Concert{}, "concert")
    defer dbmap.Db.Close()

    dbmap.Insert(&concert)
}

func FindAll() ([]Concert) {
    dbmap := db.InitDb(Concert{}, "concert")
    defer dbmap.Db.Close()

    var concerts []Concert
    _, err := dbmap.Select(&concerts, "SELECT * FROM concert ORDER BY name")
    common.CheckError(err)

    return concerts
}

func FindOne(id string) (Concert) {
    dbmap := db.InitDb(Concert{}, "concert")
    defer dbmap.Db.Close()

    var concert Concert
    err := dbmap.SelectOne(&concert, "SELECT * FROM concert WHERE id = ? ORDER BY name", id)
    common.CheckError(err)

    return concert
}

func FindConcertsByBand(bandId string) ([]Concert) {
    dbmap := db.InitDb(Concert{}, "concert")
    defer dbmap.Db.Close()

    var concerts []Concert
    _, err := dbmap.Select(&concerts, `
        SELECT c.*
        FROM band_concert bc
        JOIN concert c
        ON c.id = bc.concert_id
        WHERE bc.band_id = ?
        ORDER BY name`, bandId)
    common.CheckError(err)

    return concerts
}