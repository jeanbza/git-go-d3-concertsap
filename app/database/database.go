package db

import (
    "git-go-d3-concertsap/app/common"
    "github.com/coopernurse/gorp"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
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

type State struct {
    Id      int64   `db:"id"`
    Name    string  `db:"name"`
    Acronym string  `db:"acronym"`
}

func InitDb() *gorp.DbMap {
    // connect to db using standard Go database/sql API
    // use whatever database/sql driver you wish
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/concertsap")
    common.CheckError(err)

    // construct a gorp DbMap
    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

    dbmap.AddTableWithName(Concert{}, "concert").SetKeys(true, "id")
    dbmap.AddTableWithName(State{}, "state").SetKeys(true, "id")

    return dbmap
}