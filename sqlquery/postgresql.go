package main

import (
    _ "github.com/lib/pq"
    "database/sql"
    "fmt"
    "os"
)
const (
    host     = "bk64"
    port     = 5432
    user     = "fujiwara"
    password = "pass"
    dbname   = "nxscsendb"
)

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
        os.Exit(1)
    }
}

func main() {
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s " +
                            "password=%s dbname=%s sslmode=disable",
                            host, port, user, password, dbname)
    fmt.Println(psqlInfo)
    db, err := sql.Open("postgres", psqlInfo)
    checkError(err)
    // select
    var name string = "0734331001"
    rows, err := db.Query("SELECT user_id.user_name FROM user_data WHERE user_id = $1;", name)
    //rows, err := db.Query("SELECT user_id,user_name FROM user_data;")
    checkError(err)
    column1 := ""
    column2 := ""
    for rows.Next() {
        err = rows.Scan(&column1, &column2)
        checkError(err)
        fmt.Printf("Sql data. id=[%s], name=[%s].\n", column1, column2);
    }
    defer db.Close()
}
