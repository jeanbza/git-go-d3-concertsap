package database

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "git-go-d3-concertsap/src/common"
)

func Select(query string, database string) (*sql.Rows) {
    databaseDetails := common.StrCat("root:@tcp(127.0.0.1:3306)/", database)
    db, err := sql.Open("mysql", databaseDetails)
    common.CheckError(err)

    rows, err := db.Query(query)
    common.CheckError(err)

    return rows
}