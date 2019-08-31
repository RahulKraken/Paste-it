package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// try connecting to database
	db, err := sql.Open("mysql", "pasteit:pasteit@tcp(127.0.0.1:3306)/pasteit_db")
	if err != nil {
		panic(err.Error())
	}

	err = db.Ping();
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("database connection successfull!!!")
}