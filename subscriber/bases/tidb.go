package bases

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var ADDRTIDB = "35.225.2.109:4000"
var DBTIDB = "sopes1"
var ADMINTIDB = "sopes1"
var PASSTIDB = "sopes123"

func SaveLogTidb() {
	// Open database connection
	db, err := sql.Open("mysql", ADMINTIDB+":"+PASSTIDB+"-@tcp("+ADDRTIDB+")/"+DBTIDB)

	if err != nil {
		panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
	}

	defer db.Close()

	sql := "INSERT INTO cities(name, population) VALUES ('Moscow', 12506000)"
	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("The last inserted row id: %d\n", lastId)
}
