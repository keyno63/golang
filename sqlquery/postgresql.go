package sqlquery

import (
	"database/sql"
	"fmt"
	"os"

	// PostgreSQL driver
	_ "github.com/lib/pq"
)

const (
	host     = "unagipai"
	port     = 5432
	user     = "postgres"
	password = "pass"
	dbname   = "redmine"
)

// private

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "fatal: error: %s", err.Error())
		os.Exit(1)
	}
}

func selectDB(db *sql.DB) {
	var uNumber string
	uNumber = "0734331001"
	rows, err := db.Query("SELECT id, login FROM users WHERE id = $1;", uNumber)
	//rows, err := db.Query("SELECT user_id,user_name FROM user_data;")
	checkError(err)
	column1 := ""
	column2 := ""
	for rows.Next() {
		err = rows.Scan(&column1, &column2)
		checkError(err)
		fmt.Printf("Sql data. id=[%s], name=[%s].\n", column1, column2)
	}
}

// public.

// Query exec postgresql query.
func Query() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	fmt.Println(psqlInfo)
	db, err := sql.Open("postgres", psqlInfo)
	checkError(err)
	// select
	selectDB(db)
	defer db.Close()
}
