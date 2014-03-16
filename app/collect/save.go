package collect

import (
    "net/http"
    "git-go-d3-concertsap/app/database"
    "git-go-d3-concertsap/app/common"
)

func SaveForm(rw http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    common.CheckError(err)
    form := r.Form

    var (
        columns []string
        values  []string
    )

    if form["timestamp"][0] == "" {
        form["timestamp"][0] = "NOW()"
    }

    sql := "INSERT INTO "+form["database"][0]+"("
    
    for column := range form {
        if column != "database" {
            columns = append(columns, column)
            values = append(values, form[column][0])
        }
    }

    i := 0
    for index := range columns {
        if i != 0 {
            sql += ","
        }

        sql += columns[index]
        i++
    }

    sql += ") VALUES ("
    i = 0

    for index := range values {
        if i != 0 {
            sql += ","
        }

        sql += values[index]
        i++
    }

    sql += ")"

    database.Insert(sql, "concertsap")

    http.Redirect(rw, r, "/collect/ticket", http.StatusFound)
}