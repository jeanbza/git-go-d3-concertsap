package database

import (
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
    "git-go-d3-concertsap/src/common"
)

func Select(query string) (*sql.Rows) {
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/whiteboard")
    common.CheckError(err)

    rows, err := db.Query(query)
    common.CheckError(err)

    return rows
}