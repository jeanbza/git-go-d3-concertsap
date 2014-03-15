package collect

import (
     "net/http"
     "log"
     // "git-go-d3-concertsap/src/database"
     "git-go-d3-concertsap/src/common"
)

func SaveForm(rw http.ResponseWriter, r *http.Request) {
    err := r.ParseForm()
    common.CheckError(err)
    form := r.Form

    var (
        columns []string
        values  []string
    )

    sql := "INSERT INTO "+form["database"][0]+"("
    i := 0
    
    for column := range form {
        if column != "database" {
            columns = append(columns, column)
            values = append(values, form[column][0])

            if i != 0 {
                sql += ","
            }

            sql += column
            i++
        }
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

    log.Println(sql)

    // http.Redirect(rw, req, "/", http.StatusFound)
}