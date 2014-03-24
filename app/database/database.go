package db

import (
    "database/sql"

    "git-go-d3-concertsap/app/common"

    "github.com/coopernurse/gorp"
    _ "github.com/go-sql-driver/mysql"
)

func InitDb(i interface{}, dbName string) *gorp.DbMap {
    // connect to db using standard Go database/sql API
    // use whatever database/sql driver you wish
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/concertsap")
    common.CheckError(err)

    // construct a gorp DbMap
    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

    dbmap.AddTableWithName(i, dbName).SetKeys(true, "id")

    return dbmap
}