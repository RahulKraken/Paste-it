package main

import (
	"fmt"
	"github.com/RahulKraken/Paste-it/database"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
)

// database connection
var db *sql.DB

func main() {
	fmt.Println("Hello from main!!!")
	// try connecting to database
	db, err := sql.Open("mysql", "pasteit:pasteit@tcp(127.0.0.1:3306)/pasteit_db")
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("testing db here ->", db)
	database.CreateUser(db, database.User{ID: 2, UserName: "Kraken"})
	// user := database.User{
	// 	ID: 2,
	// 	UserName: "Kraken",
	// }
	// database.CreateUser(db, user)
}