package concert

import (
    "git-go-d3-concertsap/app/database"
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

func PrintAll() {
    // log.Println("bam")

    // initialize the DbMap
    // dbmap := db.InitDb()
    // defer dbmap.Db.Close()
}

func insertConcert(concert Concert) {
    dbmap := db.InitDb(Concert{}, "concert")
    defer dbmap.Db.Close()

    dbmap.Insert(&concert)
}