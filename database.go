package main

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// User struct
type User struct {
	ID int `json:"id"`
	UserName string `json:"user_name"`
}

// Paste struct
type Paste struct {

}

// db connection var
var db *sql.DB

// CRUD application methods
// create user
func createUser(user User) {
	insert, err := db.Query("INSERT INTO user VALUES(?, ?)", user.ID, user.UserName)
	if err != nil {
		panic(err.Error())
	}
	insert.Close()
}

// get user using id
func getUser(id int) User {
	data, err := db.Query("SELECT * FROM user WHERE id = ? LIMIT 1", id)
	if err != nil {
		panic(err.Error())
	}

	var user User
	err = data.Scan(&user.ID, &user.UserName)
	if err != nil {
		panic(err.Error())
	}

	return user
}

// update user using id
func updateUser(user User) User {
	_, err := db.Query("UPDATE user WHERE id = ? SET user.user_name = ?", user.ID, user.UserName)
	if err != nil {
		panic(err.Error())
	}
	
	return user
}

// delete user using id
func deleteUser(id int) {
	_, err := db.Query("DELETE FROM user WHERE id = ?", id)
	if err != nil {
		panic(err.Error())
	}
}

func main() {
	// try connecting to database
	db, err := sql.Open("mysql", "pasteit:pasteit@tcp(127.0.0.1:3306)/pasteit_db")
	if err != nil {
		panic(err.Error())
	}

	// trying to ping database
	err = db.Ping();
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("database connection successfull!!!")
}