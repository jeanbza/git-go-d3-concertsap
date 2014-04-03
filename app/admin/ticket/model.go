package ticket

import (
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/common"
)

type Ticket struct {
    Id          int64   `db:"id"`
    Price       string  `db:"price"`
    ConcertId   string  `db:"concert_id"`
    RetailerId  string  `db:"retailer_id"`
    Timestamp   string  `db:"timestamp"`
}

func insertTicket(ticket Ticket) {
    dbmap := db.InitDb(Ticket{}, "ticket_record")
    defer dbmap.Db.Close()

    dbmap.Insert(&ticket)
}

func FindAll() ([]Ticket) {
    dbmap := db.InitDb(Ticket{}, "ticket_record")
    defer dbmap.Db.Close()

    var tickets []Ticket
    _, err := dbmap.Select(&tickets, "SELECT * FROM ticket_records ORDER BY timestamp")
    common.CheckError(err)

    return tickets
}