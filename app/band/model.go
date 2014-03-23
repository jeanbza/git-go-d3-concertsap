package band

import (
    "log"
    "git-go-d3-concertsap/app/database"
)

func PrintAll() {
    log.Println("bam")

    // initialize the DbMap
    dbmap := db.InitDb()
    defer dbmap.Db.Close()
}