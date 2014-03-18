package state

import (
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/common"
)

func FindAll() []db.State {
    // initialize the DbMap
    dbmap := db.InitDb()
    defer dbmap.Db.Close()

    var states []db.State
    _, err := dbmap.Select(&states, "SELECT * FROM state ORDER BY name")
    common.CheckError(err)

    for _,state := range states {
        states = append(states, state)
    }

    return states
}