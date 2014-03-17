package concert

import (
    "log"
    "github.com/coopernurse/gorp"
    _ "github.com/go-sql-driver/mysql"
    "database/sql"
)

func PrintAll() {
    log.Println("bam")

    // initialize the DbMap
    dbmap := initDb()
    defer dbmap.Db.Close()

    // delete any existing rows
    // err := dbmap.TruncateTables()
    // checkErr(err, "TruncateTables failed")

    // create two posts
    p1 := newConcert("test", "123 test street", 1, "www.test.com")
    p2 := newConcert("bam", "123 bam street", 1, "www.bam.com")

    // insert rows - auto increment PKs will be set properly after the insert
    err := dbmap.Insert(&p1, &p2)
    checkErr(err, "Insert failed")

    // use convenience SelectInt
    count, err := dbmap.SelectInt("select count(*) from concert")
    checkErr(err, "select count(*) failed")
    log.Println("Rows after inserting:", count)

    // update a row
    p2.Name = "BOOM"
    count, err = dbmap.Update(&p2)
    checkErr(err, "Update failed")
    log.Println("Rows updated:", count)

    // fetch one row - note use of "post_id" instead of "Id" since column is aliased
    //
    // Postgres users should use $1 instead of ? placeholders
    // See 'Known Issues' below
    //
    err = dbmap.SelectOne(&p2, "select * from concert where id=?", p2.Id)
    checkErr(err, "SelectOne failed")
    log.Println("p2 row:", p2)

    // fetch all rows
    var concerts []Concert
    _, err = dbmap.Select(&concerts, "select * from concert order by id")
    checkErr(err, "Select failed")
    log.Println("All rows:")
    for x, p := range concerts {
        log.Printf("    %d: %v\n", x, p)
    }

    // delete row by PK
    count, err = dbmap.Delete(&p1)
    checkErr(err, "Delete failed")
    log.Println("Rows deleted:", count)

    // delete row manually via Exec
    _, err = dbmap.Exec("delete from concert where id=?", p2.Id)
    checkErr(err, "Exec failed")

    // confirm count is zero
    count, err = dbmap.SelectInt("select count(*) from concert")
    checkErr(err, "select count(*) failed")
    log.Println("Row count - should be zero:", count)

    log.Println("Done!")
}

type Concert struct {
    // db tag lets you specify the column name if it differs from the struct field
    Id      int64   `db:"id"`
    Name    string  `db:"name"`
    Address string  `db:"address"`
    StateId int64   `db:"state_id"`
    Website string  `db:"website"`
    Start   string  `db:"start"`
    End     string  `db:"end"`
}

func newConcert(name string, address string, stateId int64, website string) Concert {
    return Concert{
        Name: name,
        Address: address,
        StateId: stateId,
        Website: website,
        Start: "NOW()",
        End: "NOW()",
    }
}

func initDb() *gorp.DbMap {
    // connect to db using standard Go database/sql API
    // use whatever database/sql driver you wish
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/concertsap")
    checkErr(err, "sql.Open failed")

    // construct a gorp DbMap
    dbmap := &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{}}

    dbmap.AddTableWithName(Concert{}, "concert").SetKeys(true, "id")

    // add a table, setting the table name to 'posts' and
    // specifying that the Id property is an auto incrementing PK
    // dbmap.AddTableWithName(Concert{}, "concert").SetKeys(true, "Id")

    // create the table. in a production system you'd generally
    // use a migration tool, or create the tables via scripts
    // err = dbmap.CreateTablesIfNotExists()
    // checkErr(err, "Create tables failed")

    return dbmap
}

func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}