package state

import (
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/common"
)

type State struct {
    Id      int64   `db:"id"`
    Name    string  `db:"name"`
    Acronym string  `db:"acronym"`
}

func FindAll() []State {
    dbmap := db.InitDb(State{}, "state")
    defer dbmap.Db.Close()

    var states []State
    _, err := dbmap.Select(&states, "SELECT * FROM state ORDER BY acronym")
    common.CheckError(err)

    return states
}

func insertState (state State) {
    dbmap := db.InitDb(State{}, "state")
    defer dbmap.Db.Close()

    dbmap.Insert(&state) 
}